test:
	@echo "WITH EU REGION" && echo;
	export AWS_REGION=eu-west-1; echo $$AWS_REGION && go run main.go;
	@echo "----\n";
	@echo "WITH NO REGION" && echo;
	export AWS_REGION=us-east-1; echo $$AWS_REGION && go run main.go;
