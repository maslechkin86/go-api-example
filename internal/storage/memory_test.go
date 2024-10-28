package storage

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemoryStorage", func() {

	Describe("Create", func() {
		It("should create a new user", func() {
			// Arrange
			name := fmt.Sprintf("user_name_%s", uuid.NewString())
			s := NewMemoryStorage()

			// Act
			res, err := s.Create(name)

			// Assert
			Expect(err).To(BeNil())
			Expect(res.Name).To(Equal(name))
		})
	})
})
