package main

import (
	"errors"
	"fmt"
	"time"
)

type Operator interface {
	Balance() int64
	Withdraw(amount int64) error
	Deposit(amount int64) error
	Transactions() []Tx
}

type ActionKind string

const (
	ActionKindIncr ActionKind = "+"
	ActionKindDecr ActionKind = "-"
	DateFormat                = "2006-01-01 15:04:00"
)

type Tx struct {
	value     int64      // значение на которое изменилось
	action    ActionKind // действие, прибавляем или отнимаем
	createdAt time.Time
}

// Нужно вывести данные транзакции в формате сумма: +-value, time: время создания транзакции
func (t Tx) Print() string {
	return fmt.Sprintf("sum: %s%d time:%s", t.action, t.value, t.createdAt.Format(DateFormat))
}

var _ Operator = (*txOperator)(nil)

type txOperator struct {
	balance      int64
	transactions []Tx
}

func (t *txOperator) Balance() int64 {
	return t.balance
}

func (t *txOperator) Withdraw(amount int64) error {
	if amount > t.balance {
		return errors.New("не достаточно средств")
	}
	t.balance = -amount
	t.transactions = append(t.transactions, Tx{value: amount, action: ActionKindDecr, createdAt: time.Now()})
	return nil
}

func (t *txOperator) Deposit(amount int64) error {
	t.balance += amount
	t.transactions = append(t.transactions, Tx{value: amount, action: ActionKindIncr, createdAt: time.Now()})
	return nil
}

func (t *txOperator) Transactions() []Tx {
	return t.transactions
}

func main() {
	var op Operator = &txOperator{}
	_ = op
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}
	fmt.Println("Текущий баланс", op.Balance())
	time.Sleep(1 * time.Second)
	op.Deposit(100)
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}

	fmt.Println("Текущий баланс", op.Balance())
	for _, tx := range op.Transactions() {
		fmt.Println(tx.Print())
	}

}
