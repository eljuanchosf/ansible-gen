package ansibleGen

func yamlTemplate() string {
	return `---
# Add your data here
`
}

func variablesTemplate() string {
	return `# Add your varialbes here
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
