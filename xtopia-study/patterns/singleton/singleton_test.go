package singleton

import "testing"

func TestGetInstance(t *testing.T) {
    counter1 := GetInstance()

    if counter1 == nil {
        t.Error("excepted pointer to Singleton after calling GetInstance(), not nil")
    }

    exceptedCounter := counter1

    currentCount := counter1.AddOne()
    if currentCount != 1 {
        t.Errorf("After calling for the first time to count, the count must be 1 but it is %d", currentCount)
    }

    counter2 := GetInstance()
    if counter2 != exceptedCounter {
        t.Error("Excepted same instance in counter2 but it got a different instance")
    }

    currentCount = counter2.AddOne()
    if currentCount != 2 {
        t.Errorf("After calling 'AddOne' using the second counter, the current counte must be 2 but it is %d", currentCount)
    }
}
