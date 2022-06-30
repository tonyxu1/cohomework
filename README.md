## Use case and requirement
In finance, it's common for accounts to have so-called "velocity limits". In this task, you'll write a program that accepts or declines attempts to load funds into customers' accounts in real-time.

Each attempt to load funds will come as a single-line JSON payload, structured as follow:

```json
{
  "id": "1234",
  "customer_id": "1234",
  "load_amount": "$123.45",
  "time": "2018-01-01T00:00:00Z"
}
```

Each customer is subject to three limits:

- A maximum of $5,000 can be loaded per day
- A maximum of $20,000 can be loaded per week
- A maximum of 3 loads can be performed per day, regardless of amount

As such, a user attempting to load $3,000 twice in one day would be declined on the second attempt, as would a user attempting to load $400 four times in a day.

For each load attempt, you should return a JSON response indicating whether the fund load was accepted based on the user's activity, with the structure:

```json
{ "id": "1234", "customer_id": "1234", "accepted": true }
```

You can assume that the input arrives in ascending chronological order and that if a load ID is observed more than once for a particular user, all but the first instance can be ignored. Each day is considered to end at midnight UTC, and weeks start on Monday (i.e. one second after 23:59:59 on Sunday).

Your program should process lines from `input.txt` and return output in the format specified above, either to standard output or a file. Expected output given our input data can be found in `output.txt`.

You're welcome to write your program in a general-purpose language of your choosing, but as we use Go on the back-end and TypeScript on the front-end, we do have a preference towards solutions written in Go (back-end) and TypeScript (front-end).

We value well-structured, self-documenting code with sensible test coverage. Descriptive function and variable names are appreciated, as is isolating your business logic from the rest of your code.
## Solution
### Configuration
- All configuration is in `config.json` file
- Daily max load amount
- Daily max load count
- Weekly max load amount
- Output file name  
```
  if output file name is empty string (""), output will be directed to console 
```
### Assumption
1. Entries in `input.txt` are in ascending order by load_time
2. All entry in `input.txt` are valid,  and load amount is always with `$` before the decimal number.
3. Rejected `load_id` is still count for duplication:
```
  load_id: "6928" and customer_id: "562"
``` 
4. User who runs the application have sufficient privilege to write output file

### Unit test
1. All functions passed unit test
2. Purposely ignored func main() unit testing 

### Cost
1. 5 hours in total coding: 4 hours development and 1 hour unit testing
2. 2 Cups coffee 
3. A box of orange

