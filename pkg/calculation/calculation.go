package calculation

import (
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
    // Удаляем пробелы из выражения
    expression = strings.ReplaceAll(expression, " ", "")

    // Проверка на пустоту
    if expression == "" {
        return 0, ErrEmptyExpression
    }

    // Проверка на правильность скобок
    if !isValidParentheses(expression) {
        return 0, ErrIncorrectBracketPlacement
    }

    // Преобразование выражения в RPN
    rpn, err := infixToRPN(expression)
    if err != nil {
        return 0, err
    }

    // Вычисление выражения в RPN
    result, err := calculateRPN(rpn)
    if err != nil {
        return 0, err
    }

    return result, nil
}

func isValidParentheses(expression string) bool {
    stack := []rune{}
    for _, char := range expression {
        switch char {
        case '(':
            stack = append(stack, char)
        case ')':
            if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

func infixToRPN(expression string) ([]string, error) {
    rpn := []string{}
    operatorStack := []string{}
    numberBuffer := strings.Builder{}

    for _, char := range expression {
        switch {
        case isDigit(char) || char == '.':
            numberBuffer.WriteRune(char)
        case char == '+' || char == '-' || char == '*' || char == '/':
            if numberBuffer.Len() > 0 {
                rpn = append(rpn, numberBuffer.String())
                numberBuffer.Reset()
            }
            for len(operatorStack) > 0 && getPrecedence(string(char)) <= getPrecedence(operatorStack[len(operatorStack)-1]) {
                rpn = append(rpn, operatorStack[len(operatorStack)-1])
                operatorStack = operatorStack[:len(operatorStack)-1]
            }
            operatorStack = append(operatorStack, string(char))
        case char == '(':
            operatorStack = append(operatorStack, string(char))
        case char == ')':
            if numberBuffer.Len() > 0 {
                rpn = append(rpn, numberBuffer.String())
                numberBuffer.Reset()
            }
            for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
                rpn = append(rpn, operatorStack[len(operatorStack)-1])
                operatorStack = operatorStack[:len(operatorStack)-1]
            }
            if len(operatorStack) == 0 {
                return nil, ErrBracketMismatch
            }
            operatorStack = operatorStack[:len(operatorStack)-1]
        default:
            return nil, ErrInvalidCharacter
        }
    }

    if numberBuffer.Len() > 0 {
        rpn = append(rpn, numberBuffer.String())
    }

    for len(operatorStack) > 0 {
        rpn = append(rpn, operatorStack[len(operatorStack)-1])
        operatorStack = operatorStack[:len(operatorStack)-1]
    }

    return rpn, nil
}

func calculateRPN(rpn []string) (float64, error) {
    stack := []float64{}
    for _, token := range rpn {
        if isDigit(rune(token[0])) {
            num, err := strconv.ParseFloat(token, 64)
            if err != nil {
                return 0, err
            }
            stack = append(stack, num)
        } else {
            if len(stack) < 2 {
                return 0, ErrNotEnoughOperands
            }
            operand2 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            operand1 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result, err := calculateOperation(operand1, operand2, token)
            if err != nil {
                return 0, err
            }
            stack = append(stack, result)
        }
    }
    if len(stack) != 1 {
        return 0, ErrIncorrectExpression
    }
    return stack[0], nil
}

func calculateOperation(operand1, operand2 float64, operator string) (float64, error) {
    switch operator {
    case "+":
        return operand1 + operand2, nil
    case "-":
        return operand1 - operand2, nil
    case "*":
        return operand1 * operand2, nil
    case "/":
        if operand2 == 0 {
            return 0, ErrDivisionByZero
        }
        return operand1 / operand2, nil
    default:
        return 0, ErrUnknownOperation
    }
}

func isDigit(char rune) bool {
    return char >= '0' && char <= '9'
}

func getPrecedence(operator string) int {
    switch operator {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    default:
        return 0
    }
}
