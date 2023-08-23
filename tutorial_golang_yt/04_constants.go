// Go suporta _constants_ de caractere, sequência, booleano
// e valores numéricos.
package main

import (
	"fmt"
	"math"
)

// `const` declara um valor constante.
const s string = "constant"

func main() {
	fmt.Println(s)
    // Uma declaração `const` pode aparecer em qualquer lugar que uma 
    // declaração `var` possa.
	const n = 500000000
    // Expressões constantes executam aritmética com
    // precisão arbitrária.
	const d = 3e20 / n
	fmt.Println(d)
    // Uma constante numérica não tem tipo até que seja fornecido
    // um, como por uma conversão explícita.
	fmt.Println(int64(d))
    // Um número pode receber um tipo usando o
    // contexto, como uma atribuição a uma variável
    // ou chamada de função. Por exemplo,
    // `math.Sin` espera um` float64`.
	fmt.Println(math.Sin(n))
}