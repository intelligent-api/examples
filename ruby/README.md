# Ruby examples for api.intelligent-api.com

| endpoint                                                   | example folder              | required scopes                   |
| ---------------------------------------------------------- | --------------------------- | --------------------------------- |
| https://api.intelligent-api.com/v1/language/detect         | detect-language             | language.detect                   |
| https://api.intelligent-api.com/v1/pii/detect              | detect-pii                  | pii.detect                        |
| https://api.intelligent-api.com/v1/text/sentiment          | detect-sentiment            | text.sentiment                    |
| https://api.intelligent-api.com/v1/document/expenses       | extract-expense             | document.expense                  |
| https://api.intelligent-api.com/v1/document/identity       | extract-identity            | document.identity                 |
| https://api.intelligent-api.com/v1/document/text           | extract-text                | document.text                     |
| https://api.intelligent-api.com/v1/document/text/summarize | extract-text-summarize-text | document.text<br />text.summarize |
| https://api.intelligent-api.com/v1/pii/redact              | redact-pii                  | pii.redact                        |
| https://api.intelligent-api.com/v1/text/summarize          | summarize-text              | text.summarize                    |
| https://api.intelligent-api.com/v1/language/translate      | translate                   | language.translate                |

## Pre-requisites

1. Ruby 3 (https://www.ruby-lang.org/en/documentation/installation/) installed

## Usage

1. Replace all the relevant variables in the chosen Ruby project `basic-auth.rb` or `token-auth.rb` files i.e. find all the variables surrounded with `[[variable]]` and replace with a relevant value.
2. Run the project by executing the following command in the respective project folder

```shell
# for basic auth exmaple
ruby basic-auth.rb

# for token auth exmaple
ruby token-auth.rb
```
