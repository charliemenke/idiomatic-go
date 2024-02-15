package main

import "fmt"

// wrapper: decorator patter
type Pizza interface {
    GetPrice() int
}

type BasePizza struct {}
func (bp *BasePizza) GetPrice() int {
    return 10
}

type CheesePizza struct {
    pizza Pizza
}
func (cp *CheesePizza) GetPrice() int {
    price := cp.pizza.GetPrice()
    return price + 5
}

type MeatPizza struct {
    pizza Pizza
}
func (cp *MeatPizza) GetPrice() int {
    price := cp.pizza.GetPrice()
    return price + 8
}


type SaladPriceHandler interface {
    GetPrice() int
}
func GetSaladPrice(salad SaladPriceHandler) {
    fmt.Printf("price: %d\n", salad.GetPrice())
}

type SaladPriceHandlerFunc func() int
func (ph SaladPriceHandlerFunc) GetPrice() int {
   return ph() 
}


type SaladPriceDecorator func() SaladPriceHandlerFunc
func CreateSalad(decors ...SaladPriceDecorator) SaladPriceHandlerFunc {
    for _, h := range decors {
        
    } 
}
func WithCheese()  {
    return func() {
        return 2
    }
}



func SaladWithCheese() int {
    return 10
}
func SaladWithCheeseAndDressing() int {
    return 15
}
func SaladWithCheeseAndChicken() int {
    return 16
}


func main() {
    // wrapper
    pizza := &BasePizza{}

    pWithCheese := &CheesePizza{
        pizza: pizza,
    }
    pWithCheeseAndMeat := &MeatPizza{
        pizza: pWithCheese,
    }
    fmt.Printf("price is: %d\n", pWithCheeseAndMeat.GetPrice())

    // functions as bridge
    GetSaladPrice(SaladPriceHandlerFunc(SaladWithCheese))    
    GetSaladPrice(SaladPriceHandlerFunc(SaladWithCheeseAndDressing))    
    GetSaladPrice(SaladPriceHandlerFunc(SaladWithCheeseAndChicken))    
}
