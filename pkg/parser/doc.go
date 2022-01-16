package parser

/*
Pratt parse main idea association of parsing function with token types
once a token type is encountered, the parsing functions are called to parse the appropriate expression and return an AST node
each token type can have up to 2 parsing functions associated with it, depending on whether the token is founc in a prefix or an infix position

- letStatement
- returnStatement
- statement
 */
