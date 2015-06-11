%{

package interpreter

import (
 //"fmt"
)

%}

%token GOTO
%token GO
%token TO
%token SUB
%token INTEGER
%token FLOAT
%token NEWLINE
%token PRINT
%token INPUT
%token IF
%token THEN
%token FOR
%token STEP
%token NEXT
%token REM
%token DIM
%token DEF
%token WHITESPACE
%token COLON
%token COMMA
%token DOT
%token SEMICOLON
%token LINESTART
%token STRING_LITERAL
%token BOOLEAN
%token NULL
%token IDENTIFIER
%token RETURN
%token END
%token REM
%token PLUS
%token MINUS
%token LET
%token CIRCUMFLEX
%token END

%token EQUALS
%token LESS_THAN
%token GREATER_THAN
%token NOT_EQUALS

%token LEFT_PARENTHESIS
%token LEFT_BRACKET
%token LEFT_BRACE
%token RIGHT_BRACKET
%token RIGHT_BRACE
%token RIGHT_PARENTHESIS

%type	<item>
      INTEGER
      FLOAT
      STRING_LITERAL
      IDENTIFIER
      NULL
      BOOLEAN

%type	<item>
      Lines
      Line
      LineNumber
      Statement
      Statements
      GotoStatement
      EndStatement
      PrintStatement
      ForStatement
      NextStatement
      AssignmentStatement
      ReturnStatement
      GoSubStatement
      DefStatement
      IfStatement
      RelationExpression
      OperationStatement
      RemarkStatement
      Comparator
      Args
      Arg
      Symbol
      MapPropertyExpression
      AdditiveExpression
      MultiplicationExpression
      DivisionExpression
      PowerExpression
      Expression
      UnaryExpression
      Identifier
      StringLiteral
      Integer
      Float
      Null
      Boolean
      Array
      Map
      MapPropertyList
      MapProperty
      DimStatement
      FunctionCall
      ArrayDefinition
      ArrayDefinitions


%left MINUS PLUS
%left TIMES DIVIDE
%right CIRCUMFLEX LEFT_PARENTHESIS LEFT_BRACKET
%left DOT

%start program

%%

program : Lines
  {
    if v, ok := Parserlex.(*ParserLex); ok {
    	v.cb($1.(Lines))
    }
  }

Lines : Line
      {
        $$ = Lines{$1.(Line)}
      }
      | Lines Line
      {
        $$ = append($1.(Lines), $2.(Line))
      }

Line : LineNumber Statements NEWLINE
     {
       $$ = Line{Number: $1.(LineNumber), Statements: $2.(Statements)};
     }

LineNumber : INTEGER
           {
              $$ = LineNumber($1.(int64))
           }

Statements: Statement
          {
            $$ = Statements{$1.(Statement)}
          }
          | Statements COLON Statement
          {
            $$ = append($1.(Statements), $3.(Statement))
          }

Statement: GotoStatement
         | PrintStatement
         | AssignmentStatement
         | IfStatement
         | DimStatement
         | ForStatement
         | NextStatement
         | GoSubStatement
         | ReturnStatement
         | DefStatement
         | EndStatement
         | RemarkStatement
         | OperationStatement /* Catch everything else */
         ;

OperationStatement: IDENTIFIER Args
                  {
                    $$ = &OperationStatement{Operation: $1.(string), Args: $2.(Args)}
                  }

RemarkStatement: REM STRING_LITERAL
               {
                  $$ = &RemarkStatement{}
               }

EndStatement: END
            {
              $$ = &EndStatement{}
            }

DimStatement: DIM ArrayDefinitions
            {
              $$ = &DimStatement{ArrayDefinitions:$2.([]*ArrayDefinition)}
            }

DefStatement: DEF Identifier EQUALS Expression
            {
              $$ = &DefStatement{
                Identifier: $2.(*IdentifierSymbol),
                Expression: $4.(Expression),
              }
            }
            | DEF Identifier LEFT_PARENTHESIS Identifier RIGHT_PARENTHESIS EQUALS Expression
            {

              parameter := $4.(*IdentifierSymbol)
              $$ = &DefStatement{
                Identifier: $2.(*IdentifierSymbol),
                Parameter: parameter.Name,
                Expression: $7.(Expression),
              }
            }

GotoStatement: GOTO Expression
            {
              $$ = &GotoStatement{LineNumber: $2.(Expression)}
            }
            | GO TO Expression
            {
              $$ = &GotoStatement{LineNumber: $3.(Expression)}
            }

GoSubStatement: GO SUB Expression
              {
                $$ = &GoSubStatement{LineNumber: $3.(Expression)}
              }

ReturnStatement: RETURN
              {
                $$ = &ReturnStatement{}
              }

PrintStatement: PRINT Args
            {
              $$ = &PrintStatement{Args: $2.(Args)}
            }

AssignmentStatement: LET Symbol EQUALS Expression
            {
              $$ = &AssignmentStatement{Symbol: $2.(Symbol), Expression: $4.(Expression)}
            }
            | LET FunctionCall EQUALS Expression /* WTF? Array access is done with the same syntax as function calls */
            {
              $$ = &AssignmentStatement{Symbol: $2.(Symbol), Expression: $4.(Expression)}
            }
            | LET MapPropertyExpression EQUALS Expression
            {
              $$ = &AssignmentStatement{Symbol: $2.(Symbol), Expression: $4.(Expression)}
            }

IfStatement: IF RelationExpression THEN Expression
           {
              $$ = &IfStatement{RelationExpression: $2.(*RelationExpression), LineNumber: $4.(Expression)}
           }

ForStatement: FOR Identifier EQUALS Expression TO Expression
            {
              $$ = &ForStatement{
                ControlVariable: $2.(*IdentifierSymbol),
                InitialValue: $4.(Expression),
                Limit: $6.(Expression),
              }
            }
            | FOR Identifier EQUALS Expression TO Expression STEP Expression
            {
              $$ = &ForStatement{
                ControlVariable: $2.(*IdentifierSymbol),
                InitialValue: $4.(Expression),
                Limit: $6.(Expression),
                Increment: $8.(Expression),
              }
            }

NextStatement: NEXT Identifier
             {
               $$ = &NextStatement{
                 ControlVariable: $2.(*IdentifierSymbol),
               }
             }

RelationExpression: Expression Comparator Expression
                  {
                    $$ = &RelationExpression{Left: $1.(Expression), Comparator: $2.(Comparator), Right: $3.(Expression)}
                  }

Comparator: EQUALS
          {
            $$ = &EqualsComparator{}
          }
          | NOT_EQUALS
          {
            $$ = &EqualsComparator{Negation: true}
          }
          | LESS_THAN
          {
            $$ = &LessThanComparator{}
          }
          | LESS_THAN EQUALS
          {
            $$ = &LessThanComparator{OrEqual: true}
          }
          | GREATER_THAN
          {
           $$ = &GreaterThanComparator{}
          }
          | GREATER_THAN EQUALS
          {
            $$ = &GreaterThanComparator{OrEqual: true}
          }

Symbol: Identifier

FunctionCall: Expression LEFT_PARENTHESIS Args RIGHT_PARENTHESIS
           {
             $$ = &FunctionCall{Expression: $1.(Expression), Args: $3.(Args)}
           }

ArrayDefinitions: ArrayDefinition
                {
                  $$ = []*ArrayDefinition{$1.(*ArrayDefinition)}
                }

ArrayDefinition: Identifier LEFT_PARENTHESIS Args RIGHT_PARENTHESIS
                {
                  $$ = &ArrayDefinition{Symbol: $1.(*IdentifierSymbol), Dimensions: $3.(Args)}
                }

AdditiveExpression: Expression PLUS Expression
                  {
                    $$ = &AdditiveExpression{Left: $1.(Expression), Minus: false, Right: $3.(Expression)}
                  }
                  | Expression MINUS Expression
                  {
                    $$ = &AdditiveExpression{Left: $1.(Expression), Minus: true, Right: $3.(Expression)}
                  }

MultiplicationExpression: Expression TIMES Expression
                        {
                          $$ = &MultiplicationExpression{Left: $1.(Expression), Right: $3.(Expression)}
                        }

DivisionExpression: Expression DIVIDE Expression
                  {
                    $$ = &DivisionExpression{Left: $1.(Expression), Right: $3.(Expression)}
                  }

PowerExpression: Expression CIRCUMFLEX Expression
                 {
                   $$ = &PowerExpression{BaseExpression: $1.(Expression), ExponentExpression: $3.(Expression)}
                 }

Expression: LEFT_PARENTHESIS Expression RIGHT_PARENTHESIS
          {
            $$ = $2
          }
          | AdditiveExpression
          | MultiplicationExpression
          | DivisionExpression
          | PowerExpression
          | MapPropertyExpression
          | Symbol
          | StringLiteral
          | Integer
          | Float
          | Boolean
          | Null
          | Map
          | Array
          | FunctionCall
          | UnaryExpression

UnaryExpression: MINUS Expression
               {
                 $$ = &UnaryMinusExpression{Expression: $2.(Expression)}
               }

Identifier: IDENTIFIER
      {
        $$ = &IdentifierSymbol{Name: $1.(string)}
      }

StringLiteral: STRING_LITERAL
      {
        str := $1.(string)
        $$ = &StringLiteral{Value: str[1:len(str)-1]}
      }

Integer: INTEGER
       {
         $$ = &Integer{Value: $1.(int64)}
       }

Float: FLOAT
     {
       $$ = &Float{Value: $1.(float64)}
     }

Boolean: BOOLEAN
       {
         $$ = &Boolean{Value: $1.(bool)}
       }

Null: NULL
    {
      $$ = &Null{}
    }

Array: LEFT_BRACKET Args RIGHT_BRACKET
     {
       $$ = &Array{Items: $2.(Args)}
     }

Map: LEFT_BRACE MapPropertyList RIGHT_BRACE
   {
      $$ = &Map{Properties: $2.([]*MapProperty)}
   }

MapPropertyExpression: Expression DOT IDENTIFIER
                 {
                    $$ = &MapPropertyExpression{Map: $1.(Expression), Property: &StringLiteral{Value: $3.(string)}}
                 }
                 | Expression LEFT_BRACKET Expression RIGHT_BRACKET
                 {
                    $$ = &MapPropertyExpression{Map: $1.(Expression), Property: $3.(Expression)}
                 }

MapPropertyList: /* Empty */
       {
         $$ = []*MapProperty{}
       }
       | MapProperty
       {
         $$ = []*MapProperty{$1.(*MapProperty)}
       }
       | MapPropertyList COMMA MapProperty
       {
         $$ = append($1.([]*MapProperty), $3.(*MapProperty))
       }

MapProperty: IDENTIFIER COLON Expression
           {
             $$ = &MapProperty{Name: $1.(string), Value: $3.(Expression)}
           }

Args: /* Empty */
    {
      $$ = Args{}
    }
    | Arg
    {
      $$ = Args{$1.(Expression)}
    }
    | Args Separator Arg
    {
      $$ = append($1.(Args), $3.(Expression))
    }

Separator: COMMA
         | SEMICOLON

Arg: Expression

%%
