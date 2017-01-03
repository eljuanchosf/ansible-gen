package ansibleGen

func yamlTemplate() string {
	return `---
# Add your data here
`
}

func variablesTemplate() string {
	return `# Add your variables here
`
}

func inventoryTemplate() string {
	return `[loadbalancer]
lb01

[webserver]
app01
app02
`
}
