# Example code for using the Intelligent API

## Currently Available Language Examples

| Language | Examples                                           | Notes                                                       |
| -------- | -------------------------------------------------- | ----------------------------------------------------------- |
| C#       | https://github.com/intelligent-api/examples/csharp | Implementation without addition of any nuget packages       |
| cURL     | https://github.com/intelligent-api/examples/curl   |                                                             |
| Go       | https://github.com/intelligent-api/examples/golang |                                                             |
| Java     | https://github.com/intelligent-api/examples/java   | Implementation without addition of any dependencies         |
| Node.js  | https://github.com/intelligent-api/examples/nodejs |                                                             |
| Python   | https://github.com/intelligent-api/examples/python | Implementation using standard http.client                   |
| Ruby     | https://github.com/intelligent-api/examples/ruby   |                                                             |
| Rust     | https://github.com/intelligent-api/examples/rust   | Implementation using base64, reqwest, serde, tokio packages |

## How to use

1. Select your chosen programming language and open the relevant folder with the code examples.
2. For each endpoint there is a folder with specific code to show how to invoke the endpoint using either Basic or OAuth (token) authentication.
3. To begin ensure you have signed up at https://dash.intelligent-api.com and have generated at least one set of credentials (either a Basic or OAuth set of credentials).
4. You must keep those credentials secure, and never share them or check them into a source control repository.
5. Follow the instructions in the README.md in the respective root folder of the chose programming language, ensuring you have the pre-requisites installed.
6. When you are ready to test a particular endpoint open the respective endpoint folder and then replace the variables in the relevant code files before executing the code.

### Variables to be replaced

Within the code files the following placeholder variables need to be replaced before testing the various endpoints:

| Variable Name       | Variable Placeholder  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| ------------------- | --------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Client Id           | `[[client_id]]`       | The client id for the set of credentials being used. This is sourced from the dashboard (https://dash.intelligent-api.com/platform/apiaccess)                                                                                                                                                                                                                                                                                                       |
| Client Secret       | `[[client_secret]]`   | The client secret for the set of credentials being used. This is sourced from the dashboard (https://dash.intelligent-api.com/platform/apiaccess)                                                                                                                                                                                                                                                                                                   |
| Scopes              | `[[scopes]]`          | When using OAuth (token) authentication, the initial call to the `token` endpoint will require the caller to specify what scopes to grant the token being generated access to. You can specify all the scopes available and re-use the same token for all the endpoints, or you can create a token specifically for the endpoint being called. For more information on Scopes please see https://dash.intelligent-api.com/platform/apiaccess#scopes |
| Body of Text        | `[[text_body]]`       | When working with endpoints the work with text, a body of text will need to be supplied in certain instances.                                                                                                                                                                                                                                                                                                                                       |
| Source Language     | `[[source_language]]` | For endpoints that work with text the source language of the text will need to be supplied in certain instances, if the source language is unkown it is advisable to invoke the `detect language` endpoint and retrieving the source language before calling the respective endpoint.                                                                                                                                                               |
| Target Language     | `[[target_language]]` | For endpoints that work with text the target language for the output of the endpoint will need to be supplied in certain instances.                                                                                                                                                                                                                                                                                                                 |
| Minimum Score       | `[[minimum_score]]`   | For the `redact pii` endpoint this minimum score is optional as it indicates the threshold for which possibly detected PII should be replaced. For example if the minimum score specified is 0.75 and a word scores 0.70 as the possibility for being PII, it will not be replaced but anything with a score higher than 0.75 will                                                                                                                  |
| File Name Full Path | `[[filepath]]`        | When working with endpoints that require file binary data the full path to the respective file path will need to be supplied.                                                                                                                                                                                                                                                                                                                       |
| Content Type        | `[[content_type]]`    | Specifically for the cURL endpoints where the content type can not be calculated this variable will need to be replaced manually when working with endpoints that require file binary data.                                                                                                                                                                                                                                                         |
