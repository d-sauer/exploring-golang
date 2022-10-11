package user_test

import (
	"errors"
	"github.com/d-sauer/exploring-golang/hello-interface-mock/mocks"
	"github.com/d-sauer/exploring-golang/hello-interface-mock/user"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Do to be called once with 42 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(gomock.Any(), "Hello GoMock").Return(nil).Times(1)

	gomock.InOrder(
		mockDoer.EXPECT().DoSomething(1, "first this"),
		mockDoer.EXPECT().DoSomething(2, "then this"),
		mockDoer.EXPECT().DoSomething(3, "then this"),
		mockDoer.EXPECT().DoSomething(4, "finally this"),
	)

	testUser.Use()
}

func TestUseInOrder(t *testing.T) {
	// setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// expectation
	gomock.InOrder(
		mockDoer.EXPECT().DoSomething(1, "first this"),
		mockDoer.EXPECT().DoSomething(2, "then this"),
		mockDoer.EXPECT().DoSomething(3, "then this"),
		mockDoer.EXPECT().DoSomething(4, "finally this"),
	)

	// execute
	testUser.UseInOrder(1, "first this")
	testUser.UseInOrder(2, "then this")
	testUser.UseInOrder(3, "then this")
	testUser.UseInOrder(4, "finally this")
}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("Dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expectation
	mockDoer.EXPECT().DoSomething(42, "Hello GoMock").Return(dummyError).Times(1)

	err := testUser.Use()
	if err != dummyError {
		t.Fail()
	}
}
