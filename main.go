package main

import (
	"fmt"
	"sync"
	"time"
	"math/big"
	"strconv"

	token "./contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	INCREMENT uint64= 1000
	PROJECT_ID = ""

	URL string = "https://mainnet.infura.io/v3/363e67ae20d84f858b567d4c416df486"
)

var (
	no_of_events = 0
	tokenId = big.NewInt(1083637)
	end_block uint64 = 7028323
	start_block uint64 = 6607985
	tmp_end = start_block
	tmp_start = start_block	
	contract_address string = "0x06012c8cf97BEaD5deAe237070F9587f8E7A266d"
)

type Instance struct {
	instance	token.Token
}

type Kitty struct {
	generation int
	genes      string
	birth      string
}

type CounterKitties struct {
	mu 				 sync.Mutex
	counter  map[string]int
}

var kitties *CounterKitties = &CounterKitties{
	counter: make(map[string]int),
}

func main() {

	var wg sync.WaitGroup

	client, err := ethclient.Dial(URL)
	if err != nil {
		panic(err)
	}

	token_address := common.HexToAddress(contract_address)
	instance, err := token.NewToken(token_address, client)
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println("error is", err)
	}

	diff := end_block - start_block
	mod := diff%INCREMENT
	units := uint64(diff/INCREMENT)
	
	if mod > 0 {
		units += 1
	}

	wg.Add(int(units))
	
	for i:=0; i<int(units); i++ {
		_diff := end_block - tmp_start
		if _diff > INCREMENT {
			tmp_end += INCREMENT
		} else {
			tmp_end += _diff + 1
		}
		go fetchLogs(tmp_end, tmp_start, instance, &wg, i)
		tmp_start = tmp_end + 1
		time.Sleep(50 * time.Millisecond)
	}
	wg.Wait()
	matrons := getMatron(kitties)
	fmt.Println("-------MOST FERTILE MATRONS------")
	for k,v := range matrons[0] {
		for _, item := range v {

			_key, _ := strconv.Atoi(item)
			kitty, err := instance.GetKitty(nil, big.NewInt(int64(_key))) 
			if err != nil {
				panic(err)
			}
			fmt.Printf("kitty [%s] was one of the most fertile with [%d] appearances on events and has generation [%s], birthTime [%s] and genes [%s]\n\n", 
						 item, k, kitty.Generation.String(), kitty.BirthTime.String(), kitty.Genes.String())
		}
	}
}

func fetchLogs(end uint64, start uint64, instance *token.Token, _wg *sync.WaitGroup, id int) {

	fmt.Println("Starting routine number:", id)

	_opts := &bind.FilterOpts{
		Start: start,
		End:   &end,
	}

	events, err := instance.FilterBirth(_opts)
	if err != nil {
		panic(err)
	}

	for {
		if !events.Next() {
			break
		}
		matronId := events.Event.MatronId.String()
		kitties.mu.Lock()
		assignKitty(matronId, kitties)
		kitties.mu.Unlock()
	}
	defer _wg.Done()
	fmt.Println("finishing routine: ", id)
}

func assignKitty(matronId string, _kitties *CounterKitties) {
	if _, ok := _kitties.counter[matronId]; ok {
		_kitties.counter[matronId] += 1
	} else {
		_kitties.counter[matronId] = 1
	}
}

func kittyIdToString(id big.Int) string {
	return id.String()
}

func getMatron(_kitties *CounterKitties) ([](map[int][]string)) {
	maxValues := [](map[int]([]string)){}
	stackLimit := 3
	delete(_kitties.counter, "0")
	for k, v := range _kitties.counter {
		
		appearances := v

		if len(maxValues) == 0 {
			maxValues = append(maxValues, map[int]([]string){v:[]string{k}})
		} else {
			firstKey:= getMatronAppearancesNumber(maxValues[0])
			if appearances > firstKey {
				if len(maxValues) >= stackLimit {
					maxValues = maxValues[:len(maxValues)-1]
				}
				maxValues = append([](map[int]([]string)){map[int]([]string){appearances:[]string{k}}}, maxValues...)
			} else if  appearances == firstKey {
				newKeysSlice := append(maxValues[0][appearances], k) 
				maxValues[0][appearances] = newKeysSlice
			} else {
				continue
			}
		}
	}
	return maxValues
}

func getMatronAppearancesNumber(_map map[int]([]string)) int {
	_key := -1
	for k, _ := range _map {
		_key = k
		break
	}
	return _key
}
