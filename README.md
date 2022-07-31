# Go Flow Safe: Ensure Data Flow of your App with Taint Analyses

Presented at GopherConEU 2022 :green_heart: with these [slides](https://akwickert.de/talk/talk-go-flow-safe-ensure-data-flow-of-your-app-with-taint-analyses-gophercon-europe/202207_GopherConEU.pdf).

## Examples

To test [Go Flow Leeve](https://github.com/google/go-flow-levee), I created two minimal working examples that cover potential vulnerabilities: Leaking of sensitive information (CWE-522) and SQL injections (CWE-89). 
The configuration to execute the analysis and the output I produced is available in the respective README files ([README CWE-522](./example_sensitiveinformation/README.md) and [README CWE-89](./example_sqlInjection/README.md`w
))

## More Ressources

- More examples of how to use regular expressions and combine several Sources and Sinks: <https://github.com/google/go-flow-levee/blob/master/configuration/example-config.yaml>
- Example for semgrep SQL Injection query (much more complex): <https://semgrep.dev/orgs/akwick/editor/r/go.lang.security.injection.tainted-sql-string.tainted-sql-string>