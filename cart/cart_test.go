package cart_test

import (
	"github.com/midoblgsm/shoppingcart/cart"
	"github.com/midoblgsm/shoppingcart/resources"
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
			Expect(res.TotalCost).To(Equal(float32(0)))
		})

		It("should initially has 0 items", func() {
			res := myCart.TotalItems()
			Expect(res.Error).NotTo(HaveOccurred())
			Expect(res.TotalItems).To(Equal(0))
		})

	})

	Context(".AddItem", func() {
		var (
			item resources.Item
		)
		BeforeEach(func() {
			item = resources.Item{}
		})

		It("should fail when the item is empty", func() {
			req := resources.AddItemRequest{Item: item}
			resp := myCart.AddItem(req)

			Expect(resp.Error).To(HaveOccurred())
			Expect(resp.Error.Error()).To(Equal("cannot add an empty item"))

			cost := myCart.TotalCost()
			Expect(cost.Error).NotTo(HaveOccurred())
			Expect(cost.TotalCost).To(Equal(float32(0)))

			items := myCart.TotalItems()
			Expect(items.Error).NotTo(HaveOccurred())
			Expect(items.TotalItems).To(Equal(0))

		})
		Context(".AddPartiallyEmptyItem", func() {
			BeforeEach(func() {
				item = resources.Item{
					Name:     "potato",
					Price:    1.0,
					Quantity: 1,
				}
			})
			It("should fail to add an Item without ID", func() {
				req := resources.AddItemRequest{Item: item}
				resp := myCart.AddItem(req)
				Expect(resp.Error).To(HaveOccurred())
				Expect(resp.Error.Error()).To(Equal("cannot add an item without and ID"))
			})

			It("should fail to add an Item without price", func() {
				item = resources.Item{
					ID:       "vg1",
					Name:     "potato",
					Quantity: 1,
				}
				req := resources.AddItemRequest{Item: item}
				resp := myCart.AddItem(req)
				Expect(resp.Error).To(HaveOccurred())
				Expect(resp.Error.Error()).To(Equal("cannot add an item without a price"))
			})

			It("should succeed to add an Item without a name or quantity", func() {
				item = resources.Item{
					ID:    "vg1",
					Price: 1.0,
				}
				req := resources.AddItemRequest{Item: item}
				resp := myCart.AddItem(req)
				Expect(resp.Error).NotTo(HaveOccurred())

				cost := myCart.TotalCost()
				Expect(cost.Error).NotTo(HaveOccurred())
				Expect(cost.TotalCost).To(Equal(float32(1.0)))

				items := myCart.TotalItems()
				Expect(items.Error).NotTo(HaveOccurred())
				Expect(items.TotalItems).To(Equal(1))

			})

			It("should succeed to add an Item twice and update the total and items", func() {
				item = resources.Item{
					ID:       "vg1",
					Price:    1.0,
					Quantity: 3,
				}
				cost0 := myCart.TotalCost()
				Expect(cost0.Error).NotTo(HaveOccurred())
				Expect(cost0.TotalCost).To(Equal(float32(0)))

				items0 := myCart.TotalItems()
				Expect(items0.Error).NotTo(HaveOccurred())
				Expect(items0.TotalItems).To(Equal(0))

				req := resources.AddItemRequest{Item: item}
				resp := myCart.AddItem(req)
				Expect(resp.Error).NotTo(HaveOccurred())

				cost1 := myCart.TotalCost()
				Expect(cost1.Error).NotTo(HaveOccurred())
				Expect(cost1.TotalCost).To(Equal(float32(3.0)))

				items1 := myCart.TotalItems()
				Expect(items1.Error).NotTo(HaveOccurred())
				Expect(items1.TotalItems).To(Equal(1))

				resp = myCart.AddItem(req)
				Expect(resp.Error).NotTo(HaveOccurred())

				cost2 := myCart.TotalCost()
				Expect(cost2.Error).NotTo(HaveOccurred())
				Expect(cost2.TotalCost).To(Equal(float32(6.0)))

				items2 := myCart.TotalItems()
				Expect(items2.Error).NotTo(HaveOccurred())
				Expect(items2.TotalItems).To(Equal(1))

				Expect(cost1.TotalCost).NotTo(Equal(cost2.TotalCost))
				Expect(items1.TotalItems).To(Equal(items2.TotalItems))
			})
		})
	})

	Context(".DeleteItem", func() {

		It("should not modify total or items when ID is empty", func() {
			req := resources.RemoveItemRequest{}
			resp := myCart.RemoveItem(req)

			Expect(resp.Error).NotTo(HaveOccurred())

			cost := myCart.TotalCost()
			Expect(cost.Error).NotTo(HaveOccurred())
			Expect(cost.TotalCost).To(Equal(float32(0)))

			items := myCart.TotalItems()
			Expect(items.Error).NotTo(HaveOccurred())
			Expect(items.TotalItems).To(Equal(0))

		})
	})

})

// The application that we will build is a simple shopping cart. Let's start by listing our requirements for this shopping cart:
// It should allow the user to add items,
// It should allow the user to remove items,
// It should allow the user to get the total number of unique items
// It should allow the user to retrieve the total cost at any point.
