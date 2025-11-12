# ado-api-inventory

A Go utility for inventorying Azure DevOps organizations to assist in assessments and migrations. This tool pulls and catalogs the most relevant objects from your Azure DevOps tenant and exports them to CSV format for analysis.

## Overview

`ado-api-inventory` is designed for Azure DevOps administrators, compliance officers, and assessment professionals who need to quickly understand the structure and contents of their Azure DevOps organization. The tool provides a straightforward way to export key organizational data without requiring manual portal exploration.

### What Gets Inventoried

Currently supported inventory items:

- **Users**: User accounts and identities in your organization
- **Teams**: Team definitions and organizational structure
- **Projects**: Azure DevOps projects
- **Repositories**: Git repositories across projects

Additional endpoints will be added as the tool matures.

## Status

⚠️ **Early Development**: This tool is actively being developed and may change significantly. It is not recommended for production use at this time. Use at your own risk for testing and evaluation purposes only.

## Requirements

- **Go**: 1.21 or later (tested with 1.25)
- **Azure**: An Azure App Registration with appropriate permissions to read Azure DevOps data 

## Installation

### From Source

```bash
git clone https://github.com/joltedbot/ado-api-inventory.git
cd ado-api-inventory
go build -o ado-api-inventory
```

This will create an executable `ado-api-inventory` binary in the current directory.

## Usage

### Prerequisites

Before running the tool, you must configure the following environment variables:

| Environment Variable | Description | Example |
|---|---|---|
| `ADO_TENANT_ID` | Your Azure tenant ID | `12345678-1234-1234-1234-123456789012` |
| `ADO_CLIENT_ID` | Your App Registration client ID | `87654321-4321-4321-4321-210987654321` |
| `ADO_CLIENT_SECRET` | Your App Registration client secret | `your-secret-value` |
| `ADO_ORGANIZATION` | Your Azure DevOps organization name | `my-organization` |

**Security Notes:**
- Store credentials securely; never commit them to version control
- Credentials are not persisted between runs
- Ensure your App Registration has minimal required permissions (read-only access recommended)

### Setting Up an Azure App Registration

1. In Azure Portal, navigate to Azure Active Directory → App registrations
2. Click "New registration" and provide a name
3. Under "Certificates & secrets", create a new client secret
4. Assign the necessary permissions to access Azure DevOps resources
5. Note your Tenant ID, Client ID, and Client Secret
6. Set all four environment variables (see Prerequisites section above)

### Running the Tool

Once environment variables are configured, simply run the tool:

```bash
./ado-api-inventory
```

The tool will:
1. Validate all required environment variables
2. Authenticate using your App Registration credentials
3. Query your Azure DevOps organization
4. Export inventory data to the `output/` directory as CSV files

### Output

Exported CSV files are created in the `output/` directory:

- `users.csv` - User accounts and identities
- `teams.csv` - Team definitions
- `projects.csv` - Project information
- `repositories.csv` - Repository details

## Dependencies

- [microsoft-authentication-library-for-go](https://github.com/AzureAD/microsoft-authentication-library-for-go) - Azure authentication
- [validator](https://github.com/go-playground/validator) - Data validation

## License

This project is licensed under the [Apache License 2.0](LICENSE). See the LICENSE file for details.

## Disclaimer

This tool is provided as-is for assessment and evaluation purposes. Use at your own risk. The maintainers are not responsible for any data loss or unintended consequences resulting from the use of this tool. Always test in a non-production environment first.
