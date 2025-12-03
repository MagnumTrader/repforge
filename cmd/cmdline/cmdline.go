package cmdline

import (
	"fmt"

	"github.com/MagnumTrader/repforge/internal/services"
)

func main()  {

	
	// what should we do, test the services


	db := infrastructure.NewDb()
	services.ExerciseService

	
	fmt.Println("hello world")
}
