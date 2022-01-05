package middlewares

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func Ether(x string) *big.Int {

	f , err := strconv.ParseFloat(x , 64)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	etherString := fmt.Sprintf("%f ", 1000000000000000000 * f )

	res := strings.Split(etherString, ".")
	Wei , _ := strconv.Atoi(res[0])

	return big.NewInt(int64(Wei))


}

