package cart_test

import (
	"github.com/midoblgsm/shoppingcart/cart"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cart", func() {
	var (
		myCart cart.CartInterface
	)
	BeforeEach(func() {
		myCart = cart.NewCart()
	})

	Context(".NewCart", func() {

		It("should initially has 0 total", func() {
			res := myCart.TotalCost()
			Expect(res.Error).NotTo(HaveOccurred())
		})

	})
})

// 	Context(".AddItem", func() {
// 		var (
// 			item resources.Item
// 		)
// 		BeforeEach(func() {
// 			item = resources.Item{"p1", "potato", 1.0, 1}
// 		})
//
// 		It("should add 1 item to the cart", func() {
//
// 		})
//
// 		It("should add the item's price to the total", func() {
// 		})
// 	})
//
// })

// The application that we will build is a simple shopping cart. Let's start by listing our requirements for this shopping cart:
// It should allow the user to add items,
// It should allow the user to remove items,
// It should allow the user to retrieve the total cost at any point.
