#!/bin/bash

# PROJECT_NAME Repository Initialization Script
# This script replaces placeholders in the template repository with your actual project name and owner

set -e

PROJECT_NAME="${1}"
OWNER="${2}"

if [ -z "$PROJECT_NAME" ] || [ -z "$OWNER" ]; then
    echo "Usage: $0 <PROJECT_NAME> <GITHUB_OWNER>"
    echo ""
    echo "Example: $0 my-awesome-app myusername"
    exit 1
fi

# Validate project name (basic validation)
if [[ ! "$PROJECT_NAME" =~ ^[a-z0-9-]+$ ]]; then
    echo "Error: PROJECT_NAME must contain only lowercase letters, numbers, and hyphens"
    exit 1
fi

# Convert PROJECT_NAME to various formats
PROJECT_NAME_UPPER=$(echo "$PROJECT_NAME" | tr '[:lower:]' '[:upper:]')
PROJECT_NAME_TITLE=$(echo "$PROJECT_NAME" | sed 's/-/ /g' | awk '{for(i=1;i<=NF;i++)sub(/./,toupper(substr($i,1,1)),$i)}1' | sed 's/ //g')

echo "Initializing repository..."
echo "  Project Name: $PROJECT_NAME"
echo "  GitHub Owner: $OWNER"
echo ""

# Function to replace placeholders in a file
replace_in_file() {
    local file="$1"
    if [ -f "$file" ]; then
        # Use different sed syntax for macOS vs Linux
        if [[ "$OSTYPE" == "darwin"* ]]; then
            sed -i '' "s/PROJECT_NAME/$PROJECT_NAME/g" "$file"
            sed -i '' "s/OWNER/$OWNER/g" "$file"
            sed -i '' "s/PROJECT_NAME_UPPER/$PROJECT_NAME_UPPER/g" "$file"
            sed -i '' "s/PROJECT_NAME_TITLE/$PROJECT_NAME_TITLE/g" "$file"
        else
            sed -i "s/PROJECT_NAME/$PROJECT_NAME/g" "$file"
            sed -i "s/OWNER/$OWNER/g" "$file"
            sed -i "s/PROJECT_NAME_UPPER/$PROJECT_NAME_UPPER/g" "$file"
            sed -i "s/PROJECT_NAME_TITLE/$PROJECT_NAME_TITLE/g" "$file"
        fi
    fi
}

# Replace in go.mod
echo "  Updating go.mod..."
replace_in_file "go.mod"

# Replace in README.md
echo "  Updating README.md..."
replace_in_file "README.md"

# Replace in PROJECT_STRUCTURE_GUIDE.md
echo "  Updating PROJECT_STRUCTURE_GUIDE.md..."
replace_in_file "PROJECT_STRUCTURE_GUIDE.md"

# Replace in Makefile
echo "  Updating Makefile..."
replace_in_file "Makefile"

# Replace in Dockerfile
echo "  Updating Dockerfile..."
replace_in_file "Dockerfile"

# Replace in config.yaml
echo "  Updating config.yaml..."
replace_in_file "config.yaml"

# Replace in GoReleaser configs
echo "  Updating GoReleaser configs..."
replace_in_file ".goreleaser.yaml"
replace_in_file ".goreleaser.pr.yaml"

# Replace in Go source files
echo "  Updating Go source files..."
find . -type f -name "*.go" ! -path "./vendor/*" ! -path "./.git/*" -exec grep -l "OWNER\|PROJECT_NAME" {} \; | while read -r file; do
    replace_in_file "$file"
done

# Replace in YAML files
echo "  Updating YAML files..."
find . -type f \( -name "*.yaml" -o -name "*.yml" \) ! -path "./vendor/*" ! -path "./.git/*" -exec grep -l "OWNER\|PROJECT_NAME" {} \; | while read -r file; do
    replace_in_file "$file"
done

# Replace in workflow files
echo "  Updating GitHub workflow files..."
find .github/workflows -type f \( -name "*.yaml" -o -name "*.yml" \) 2>/dev/null | while read -r file; do
    if [ -f "$file" ]; then
        replace_in_file "$file"
    fi
done

# Replace in CRD and RBAC files
echo "  Updating Kubernetes manifests..."
find config -type f \( -name "*.yaml" -o -name "*.yml" \) 2>/dev/null | while read -r file; do
    if [ -f "$file" ]; then
        replace_in_file "$file"
    fi
done

# Replace in dotfiles
echo "  Updating dotfiles..."
replace_in_file ".pre-commit-config.yaml"
replace_in_file ".yamllint.yaml"
replace_in_file "renovate.json"
replace_in_file ".vscode/tasks.json"

# Update module references in go.mod (if needed)
echo "  Verifying go.mod..."
if grep -q "github.com/OWNER/PROJECT_NAME" go.mod; then
    echo "    Warning: go.mod still contains placeholder. Please verify manually."
fi

echo ""
echo "✓ Initialization complete!"
echo ""
echo "Next steps:"
echo "  1. Review and update configuration files as needed"
echo "  2. Run: go mod tidy"
echo "  3. Run: make generate"
echo "  4. Start implementing your application!"
echo ""

