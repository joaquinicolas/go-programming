package main

import (
	"fmt"
	"gopl.io/ch7/eval"
	"math"
	"net/http"
	"strings"
)

type Expr interface {
	Eval(env Env) float64
	Check (vars map[Var]bool) error
}


type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

type literal float64

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

type unary struct {
	op rune
	x Expr
}

func (u unary) Check(vars map[Var] bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}

	return u.x.Check(vars)
}


func (u unary) Eval (env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)

	}

	panic(fmt.Sprintf("unsopported unary operator: %q", u.op))
}

type binary struct {
	op rune
	x,y Expr
}

func (b binary) Check(vars map[Var] bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}

	return b.y.Check(vars)
}


func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	}

	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

type call struct {
	fn string
	args []Expr
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}

	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}

	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}

	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}


type Env map[Var]float64

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}

	vars := make(map[Var]bool)
	if err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: " + err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/sbg+xml")
	
}

func main() {
	
}
