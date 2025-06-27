package task

import (
	"context"
	"math/rand"
	"time"
)

func Task() (int) {
r:= rand.New(rand.NewSource(time.Now().UnixNano()))
context_work := r.Intn(3) + 3 
ctx, cancel :=	context.WithTimeout(context.Background(), time.Duration(context_work)*time.Minute) // тут рандомно от 3 до 5 минут работает контекст как по условию
	defer cancel()
	go func ()  {
		<-ctx.Done()
	} ()
}
