## Notações BIG O

É a forma de definir o limite SUPERIOR (upper bound) de um algoritmo de execução.

## Exemplo

Se você disser que um algoritmo executa em no máximo 100s em seu pior tempo, isso quer dizer que temos o upper bound, onde não teríamos centários onde demoraria mais do que 100s.


## Regras da notação BIG O (O)

- Cada ação de "assign" (var int x = 5) custa 1 unidade de tempo
- Cada ação de return (return x ) custa 1 unidade de tempo
- Cada operação de aritmética (x + y) custa 1 unidade de tempo
- Cada operação lógica (x && y) custa 1 unidade de tempo
- Outras pequenas operações também levam 1 unidade de tempo

- Para outras ordem menores, sempre desconsiderar:
        - T = n² + 3n + 1 ==> O (n²)

- Desconsiderar constantes
        - T = 3n² + 6n + 1 ==> O(n²)