# go-cli-tax-calculation
GoLang playground building a CLI to calculate tax over stock trading

The CLI was built on the premise of having a input of arrays of orders as of the following pattern
``json
{ operation: buy/sell , quantity: integer, unit-cost: float }
``

## Localy
with Go 1.16 installed, you can simply build the program and run it passing -input=[] flag containing the orders

### Building
To build the program with taxcalculator name:

``go build -o taxcalculator .``

### Testing
``go test ./utils``

Since all tests were written for the utils package and the main program solely handles inputs and calls utils functions, we call the tests from utils

### Running the program
After building the program, simply run

``./taxcalculator -input=[]`` with the desired Input of orders

Example:

``
./taxcalculator -input='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}]'
``
## Docker
If you wish to use Docker just follow the next steps
### Building the image
``docker build -t taxcalculator .``

### Running Unit tests
To run unit tests on Docker, simply build the Dockerfile.test image and then run it, as follows:

``docker build -t taxcalc-test -f Dockerfile.test .``

``docker run taxcalc-test``
### Running the program
pattern:
``docker run taxcalculator -input='[]'``

usage example:
``docker run taxcalculator -input='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}]'``

