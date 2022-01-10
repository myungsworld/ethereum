cd project/contracts

#solc --bin --abi todo.sol -o ../../../go/build

cd ../../../go

abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=todo --out=gen/todo.go