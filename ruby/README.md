# Ruby examples for api.intelligent-api.com

| endpoint                                                          | example folder                   | required scopes                         |
| ----------------------------------------------------------------- | -------------------------------- | --------------------------------------- |
| https://api.intelligent-api.com/v1/language/detect                | detect-language                  | language.detect                         |
| https://api.intelligent-api.com/v1/language/detect/pii/detect     | detect-language-detect-pii       | language.detect<br />pii.detect         |
| https://api.intelligent-api.com/v1/language/detect/pii/redact     | detect-language-redact-pii       | language.detect<br />pii.redact         |
| https://api.intelligent-api.com/v1/language/detect/translate      | detect-language-translate        | language.detect<br />language.translate |
| https://api.intelligent-api.com/v1/language/detect/text/sentiment | detect-language-detect-sentiment | language.detect<br />text.sentiment     |
| https://api.intelligent-api.com/v1/pii/detect                     | detect-pii                       | pii.detect                              |
| https://api.intelligent-api.com/v1/text/sentiment                 | detect-sentiment                 | text.sentiment                          |
| https://api.intelligent-api.com/v1/document/expenses              | extract-expense                  | document.expense                        |
| https://api.intelligent-api.com/v1/document/identity              | extract-identity                 | document.identity                       |
| https://api.intelligent-api.com/v1/document/text                  | extract-text                     | document.text                           |
| https://api.intelligent-api.com/v1/document/text/summarize        | extract-text-summarize-text      | document.text<br />text.summarize       |
| https://api.intelligent-api.com/v1/pii/redact                     | redact-pii                       | pii.redact                              |
| https://api.intelligent-api.com/v1/text/summarize                 | summarize-text                   | text.summarize                          |
| https://api.intelligent-api.com/v1/language/translate             | translate                        | language.translate                      |

## Pre-requisites

1. .net 8 installed

## Usage

1. Replace all the relevant variables in the chosen c# project `Program.cs` file i.e. find all the variables surrounded with `[[variable]]` and replace with a relevant value.
2. Run the project by executing the following command in the respective project folder

```shell
dotnet run
```
