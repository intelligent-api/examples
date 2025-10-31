# C# examples for api.intelligent-api.com

| endpoint                                                   | example folder              | required scopes                      |
| ---------------------------------------------------------- | --------------------------- | ------------------------------------ |
| https://api.intelligent-api.com/v1/language/detect         | detect-language             | language.detect                      |
| https://api.intelligent-api.com/v1/pii/detect              | detect-pii                  | pii.detect                           |
| https://api.intelligent-api.com/v1/text/sentiment          | detect-sentiment            | text.sentiment                       |
| https://api.intelligent-api.com/v1/document/expenses       | extract-expense             | document.expense                     |
| https://api.intelligent-api.com/v1/document/identity       | extract-identity            | document.identity                    |
| https://api.intelligent-api.com/v1/document/text           | extract-text                | document.text                        |
| https://api.intelligent-api.com/v1/document/text/summarize | extract-text-summarize-text | document.text<br />text.summarize    |
| https://api.intelligent-api.com/v1/pii/redact              | redact-pii                  | pii.redact                           |
| https://api.intelligent-api.com/v1/scraper/markdown        | scraper-markdown            | scraper.markdown                     |
| https://api.intelligent-api.com/v1/scraper/summarize       | scraper-summarize           | scraper.markdown<br />text.summarize |
| https://api.intelligent-api.com/v1/text/summarize          | summarize-text              | text.summarize                       |
| https://api.intelligent-api.com/v1/language/translate      | translate                   | language.translate                   |

## Pre-requisites

1. .net 8 sdk (https://dotnet.microsoft.com/en-us/download/dotnet/8.0) installed

## Usage

1. Replace all the relevant variables in the chosen c# project `Program.cs` file i.e. find all the variables surrounded with `[[variable]]` and replace with a relevant value.
2. Run the project by executing the following command in the respective project folder

```shell
dotnet run
```
