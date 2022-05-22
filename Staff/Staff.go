package staff

import (
	"fmt"
	"strconv"
	beverage "tanht/Beverage"
	input "tanht/Input"
	media "tanht/Media"
	toppings "tanht/Toppings"
)

type Staff struct {
	askCustomer input.Input
	cart Cart
	device *media.Media
}

func (this *Staff) Greeting() {
	if this.cart.GetNumbers() == 0 {
		fmt.Println("Welcome to ZaloPay Coffee Shop!!")
		fmt.Println("Which beverage do you want ?")
	} else {
		this.cart.ShowBeverage()	
		fmt.Println("Do you want more beverage?")
	}
	fmt.Println("1. Bubble Tea \t 2. Coffee \t 3. Tea \t 4. No thanks")
}


func (this *Staff) MakeBeverage(selection string) (*beverage.Beverage, int) {

	b := &beverage.Beverage{}

	bubbleTea := &beverage.BubbleTea{}
	coffee := &beverage.Coffee{}
	tea := &beverage.Tea{}

	switch selection {
	case "1":
		b.SetRecipe(bubbleTea)
	case "2":
		b.SetRecipe(coffee)
	case "3":
		b.SetRecipe(tea)
	case "4":
		fmt.Println("See you again!!")
		return b, -1
	default:
		fmt.Println("Please choose correctly!!")
	}

	b.SetPrepare()

	return b, 0
}

func (this *Staff) MakeChoosen() string {

	var selection string

	for {

		selection = this.askCustomer.IReader()
		number, err := strconv.Atoi(selection)
		if err == nil {
			if number < 5 {
				break
			}
		}
	}

	return selection

}

func (this *Staff) ShowToppingsChoosing() {
	fmt.Println("1. BrownSugar \t 2. Bubble \t 3. Jelly")
}

func (this *Staff) GetInstructionTopping(e toppings.Composite) string {
	if len(e.Components) == 0 {
		fmt.Println("There are no toppings in your beverage, do you want more ?")
	} else {
		fmt.Println("There are some toppings in your beverage:")
		e.ShowComponent()
		fmt.Println("Do you want to add more toppings ?")
	}
	fmt.Println("1. Yes \t 2. No")
	return this.askCustomer.IReader()
}

func (this *Staff) AddTopings(topping *toppings.Composite) int {
	
	listTopping := [3]string{"BrownSugar", "Bubble", "Jelly"}

	index, err := strconv.Atoi(this.askCustomer.IReader())
	if err != nil {
		fmt.Println("Please input correct!!")
		return -1
	}

	if index > 3 || index < 1 {
		fmt.Println("Please input in range [1..3]")
		return -1
	}
	
	Topping, er := toppings.FactoryTopping(listTopping[index - 1])
	if er != 0 {
		fmt.Println("wtf")
		return -1
	}

	for _, each := range topping.Components {
		if each.GetName() == Topping.GetName() {
			fmt.Println(each.GetName() + " is in your beverage already")
			return -1
		}
	}
	
	topping.AddComponent(Topping)

	return 0
} 


func (this *Staff) MakeTopping(b *beverage.Beverage) {

	topping := toppings.Composite{}
	previousTopping := 0

	for {
		var err string
		if previousTopping != -1 {
			err = this.GetInstructionTopping(topping)
			if err == "2" {
				break
			}
		}

		this.ShowToppingsChoosing()

		previousTopping = this.AddTopings(&topping)
	}

	b.SetPrice(topping)
}

func (this *Staff) ShowMediaAvailable() {

	fmt.Println("Which media do you want to recieve status order ?")
	fmt.Println("1. Zalo \t 2. Messenger \t 3. None")

}

func (this *Staff) AddZaloDevice(notification *media.Media, count *int) {
	zalo, err := media.GetInstanceZalo()
	if err == 0 {
		notification.AddMedia(zalo)
		(*count) = (*count) + 1
		fmt.Println("Add zalo success")
	} else {
		fmt.Println("Zalo has been added already")
	}
}

func (this *Staff) AddMessengerDevice(notification *media.Media, count *int) {
	mess, err := media.GetInstanceMessenger()
	if err == 0 {
		notification.AddMedia(mess)
		(*count) = (*count) + 1
		fmt.Println("Add messenger success")
	} else {
		fmt.Println("Messenger has been added already")
	}
}

func (this *Staff) MakeNotication() *media.Media {

	fmt.Print("\033[H\033[2J")

	notification := &media.Media{}
	count := 0

	for {

		this.ShowMediaAvailable()

		choose := this.askCustomer.IReader()
		if choose == "1" {
			this.AddZaloDevice(notification, &count)	
		}
		if choose == "2" {
			this.AddMessengerDevice(notification, &count)	
		}
		if count == 2 || (choose == "3" && count > 0) {
			break
		}
		if count == 0 && choose == "3" {
			fmt.Println("Please choose at least one media to notify")
		}
	}

	return notification
}

func (this *Staff) StopService(err int) bool {
	if err == -1 {
		if this.cart.GetNumbers() == 0 {
			fmt.Println("Thanks, see you again!!")
		} else {
			fmt.Println("Order success")
			this.device.SendMessage("Your order is in queue, thank you !!")
		}
		return true
	}
	return false
}

func (this *Staff) Working() {

	this.askCustomer = &input.Console{}
	this.askCustomer.ISetIo()
	this.device = this.MakeNotication()

	this.cart = Cart{}

	for {

		fmt.Print("\033[H\033[2J")
		this.Greeting()
		choose := this.MakeChoosen()
		b, err := this.MakeBeverage(choose)
		if err != -1 {
			this.MakeTopping(b)
			this.cart.AddToCart(b)
		}
		
		if this.StopService(err) {
			break
		}
	}

	nameBeverages := this.cart.GetNameBeverages()
	
	for _, each := range nameBeverages {
		this.device.SendMessage("Your " + each + " is done !!")
	}

	message := "Your order is done, total price is " + strconv.Itoa(this.cart.GetPrice())
	this.device.SendMessage(message)

}