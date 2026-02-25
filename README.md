# vault-cli

A command-line interface for [Vault](https://vault-two-lovat.vercel.app) — manage your code snippets from the terminal.

## Installation
```bash
go install github.com/yuricommits/vault-cli@latest
```

Or build from source:
```bash
git clone https://github.com/yuricommits/vault-cli
cd vault-cli
go build -o vault .
sudo mv vault /usr/local/bin/
```

## Authentication

1. Go to [vault-two-lovat.vercel.app/dashboard/settings](https://vault-two-lovat.vercel.app/dashboard/settings)
2. Create a CLI access token
3. Run:
```bash
vault auth login --token <your-token>
```

To use a self-hosted instance:
```bash
vault auth login --token <your-token> --url https://your-instance.com
```

Check auth status:
```bash
vault auth status
```

Log out:
```bash
vault auth logout
```

## Commands

### List snippets
```bash
vault list
```

### Search snippets
```bash
vault search "debounce"
```

### Copy snippet to clipboard
```bash
vault copy <id>
```

### Create a new snippet
```bash
vault new
```

### Push a local file as a snippet
```bash
vault push ./utils/debounce.ts
vault push ./utils/debounce.ts --title "useDebounce hook" --description "Debounce hook for React"
```

### Delete a snippet
```bash
vault delete <id>
vault delete <id> --force
```

## Configuration

Config is stored at `~/.config/vault/config.json`.

| Key        | Description                     | Default                              |
|------------|---------------------------------|--------------------------------------|
| `token`    | CLI access token                | —                                    |
| `base_url` | Base URL of your Vault instance | `https://vault-two-lovat.vercel.app` |

## License

MIT
