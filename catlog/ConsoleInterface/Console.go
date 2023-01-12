package ConsoleInterface

import (
	"fmt"
)

func Console() {
	fmt.Println("WELCOME TO THE CONSOLE INTERFACE")
	fmt.Println("Please make your choice to perform the task")
	fmt.Printf("1.Product\n2.Category\n3.Inventory\n4.CartItem\n")
	fmt.Println("Enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}

	if choice == 1 {
		Product()
	} else if choice == 2 {
		Category()
	} else if choice == 3 {
		Inventory()
	} else if choice == 4 {
		CartItem()
	}

}
