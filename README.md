
# backstage template CLI

## Overview
The **backstage template CLI** is a powerful tool designed to manage and automate the creation, validation, and deployment of Backstage component YAML files. With a user-friendly interface and reusable modules, this CLI simplifies your workflow for Backstage projects.

## Features
- **Generate YAML Files**: Prompt users for input and generate Backstage component YAML files.
- **Validate YAML Files**: Validate the structure and fields of Backstage YAML files.
- **Reusable Modules**: Shared structures and logic for extensibility.
- **Modular Commands**: Easily add or extend commands to fit your Backstage needs.

## Commands
### `generate`
Generates a Backstage component YAML file based on user inputs.
```bash
./backstage-template generate
```

Example:
```plaintext
Enter component name: backend
Enter description: Backend service for Pomodoro Cloud-Native project
Enter GitHub project slug (e.g., user/repo): NineKama/pomodoro-cloud-native
Enter type (e.g., service): service
Enter owner: team-backend
Enter lifecycle (e.g., production): production
```

Output:
```yaml
apiVersion: backstage.io/v1alpha1
kind: Component
metadata:
  name: backend
  description: Backend service for Pomodoro Cloud-Native project
  annotations:
    github.com/project-slug: NineKama/pomodoro-cloud-native
spec:
  type: service
  owner: team-backend
  lifecycle: production
```

### `validate`
Validates a Backstage component YAML file.
```bash
./backstage-template validate component.yaml
```

Output:
```plaintext
Validation successful! The YAML file is valid.
```

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/backstage-template.git
   cd backstage-template
   ```

2. Build the CLI:
   ```bash
   go build -o backstage-template
   ```

3. Run the CLI:
   ```bash
   ./backstage-template
   ```


## Contributing
Contributions are welcome! Please submit issues or pull requests for any improvements or new features.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.
