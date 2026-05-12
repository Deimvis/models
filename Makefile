gen: gen-go

gen-go:
	nix run github:Deimvis/cg -- \
		'./utility/definitions/**' \
		--src-type 'openapi' \
		--dst-type 'go' \
		--output-basename-suffix=.cg

