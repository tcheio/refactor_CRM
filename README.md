# 📇 Mini-CRM (CLI) en Go — CHEIO THOMAS

**Mini-CRM** est un gestionnaire de contacts professionnel en ligne de commande développé en **Go**.  
Il s’agit d’un projet complet mettant en œuvre les **bonnes pratiques de développement Go**, notamment :

- Architecture modulaire avec des packages séparés  
- Injection de dépendances via des **interfaces**  
- CLI moderne basée sur **Cobra**  
- Gestion de configuration externe via **Viper**  
- Persistance de données flexible : **SQLite (GORM)**, **JSON** ou **mémoire**

---

## 🚀 Lancer le projet

Depuis la racine du projet :

```bash
go run .
```

Ou compilez un binaire exécutable :

```bash
go build -o bin/mini-crm .
./bin/mini-crm
```

---

## 📁 Structure du projet

```
refactor_crm_interface/
├── cmd/                 # Sous-commandes CLI
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── update.go
│   └── delete.go
├── internal/
│   └── storage/        # Implémentations du stockage (SQLite, JSON, mémoire)
│       ├── storage.go
│       ├── gorm_store.go
│       ├── json_store.go
│       └── memory_store.go
├── config.yaml         # Fichier de configuration de l’application
├── data/
│   ├── contacts.db     # Base de données SQLite
│   └── contacts.json   # Fichier JSON (si utilisé)
├── go.mod
└── main.go
```

---

## ⚙️ Configuration (`config.yaml`)

Le comportement de l’application est défini dans un fichier externe qui permet de choisir le mode de persistance **sans recompilation** :

```yaml
storage:
  type: gorm       # gorm | json | memory
  json_path: data/contacts.json
  db_path:   data/contacts.db
```

- **type** : mode de stockage (`gorm`, `json` ou `memory`)  
- **json_path** : chemin du fichier de stockage JSON  
- **db_path** : chemin de la base de données SQLite  

📁 **Fichiers importants :**
- Base SQLite : `data/contacts.db`  
- Fichier JSON : `data/contacts.json`  
- Configuration : `config.yaml`

---

## 📖 Fonctionnalités principales

✅ **Gestion complète des contacts (CRUD)**  
- `add` → Ajouter un contact  
- `list` → Lister les contacts  
- `update` → Modifier un contact existant  
- `delete` → Supprimer un contact

✅ **Interface CLI moderne (Cobra)**  
- Commandes claires et standardisées  
- Aide intégrée avec `--help`

✅ **Configuration externe (Viper)**  
- Modification du backend sans recompiler  
- Gestion flexible des chemins pour les fichiers

✅ **Persistance flexible (via interfaces)**  
- `gorm` : stockage dans une base SQLite locale  
- `json` : stockage simple et lisible  
- `memory` : stockage en mémoire (tests, démo)

---

## 🧪 Exemples d’utilisation

### ➕ Ajouter un contact
```bash
go run . add --name "Alice" --email "alice@example.com"
```

### 📜 Lister les contacts
```bash
go run . list
```

### ✏️ Mettre à jour un contact
```bash
go run . update --id 1 --name "Alice Cooper"
```

### 🗑️ Supprimer un contact
```bash
go run . delete --id 1
```

---

## 🧠 Concepts Go utilisés

- **Cobra** → création de commandes et sous-commandes CLI  
- **Viper** → gestion de configuration via fichiers et variables d’environnement  
- **GORM** → ORM pour la persistance SQLite  
- **Interfaces** → injection de dépendances et découplage de la logique métier  
- **Architecture modulaire** → séparation claire des responsabilités

---

## 📚 Commandes utiles

```bash
# Aide globale
go run . --help

# Aide pour une commande spécifique
go run . add --help
```

---

## 🧩 Améliorations possibles

- Exporter/importer les contacts en CSV  
- Ajouter une commande `config show`  
- Écrire des tests unitaires (`go test`)  
- Créer une API REST réutilisant la même logique métier

---

💡 **Mini-CRM** démontre comment transformer un simple programme Go en une application CLI robuste, configurable, extensible et prête pour la production.
