# GoMon - System Monitor

GoMon est un outil de surveillance du système développé en **Go**. Il permet de récupérer et d'afficher en temps réel des informations sur l'utilisation des ressources du système telles que le CPU, la mémoire, le réseau, les disques et les ports ouverts.

## 📌 Fonctionnalités

- **Surveillance du CPU** : Utilisation globale et par cœur.
- **Surveillance de la mémoire** : Mémoire totale, utilisée, libre et en pourcentage.
- **Surveillance du disque** : Espace total, utilisé, libre et en pourcentage.
- **Surveillance du réseau** : Données envoyées et reçues.
- **Ports ouverts** : Liste des ports en écoute et des processus associés.
- **Informations système** : Nom de la machine, OS, architecture, plateforme, uptime.
- **Système d'alertes** : Détection de conditions critiques (CPU > 90%, mémoire > 85%, disque < 10% libre).
- **Interface Web** : Une page HTML dynamique qui affiche les données en temps réel.

## 📂 Structure du projet

```
GoMon/
├── monitor/        # Module contenant les fonctions de récupération des données système
│   ├── cpu.go      # Récupération de l'utilisation du CPU
│   ├── memory.go   # Récupération de l'utilisation de la mémoire
│   ├── disk.go     # Récupération de l'utilisation du disque
│   ├── network.go  # Récupération des stats réseau
│   ├── system.go   # Informations système (OS, uptime, architecture...)
│   ├── ports.go    # Récupération des ports ouverts
│   ├── alerts.go   # Gestion des alertes système
│
├── routes/         # Définition des routes HTTP pour l'API
│   ├── cpu.go      # Routes pour le CPU
│   ├── memory.go   # Routes pour la mémoire
│   ├── disk.go     # Routes pour le disque
│   ├── network.go  # Routes pour le réseau
│   ├── system.go   # Routes pour les infos système
│   ├── ports.go    # Routes pour les ports
│   ├── alerts.go   # Routes pour les alertes
│   ├── base.go     # Route pour servir l'interface web
│   ├── register.go # Registre toutes les routes
│
├── utils/          # Fonctions utilitaires
│   ├── response.go # Gestion des réponses JSON
│   ├── logging.go  # Log des requêtes
│   ├── error.go    # Gestion des erreurs
│   ├── time.go     # Formatage du temps (uptime)
│
├── web/            # Fichiers web pour l'interface utilisateur
│   ├── monitor.html # Interface utilisateur dynamique
│
├── main.go         # Point d'entrée principal du serveur
│
└── README.md       # Documentation du projet
```

## 🚀 Installation et Exécution

### 📥 Prérequis
- **Go 1.18+** installé sur votre machine
- Modules Go gérés avec `go mod`

### 🔧 Installation
```sh
git clone https://github.com/MederickBernier/GoMon.git
cd GoMon
go mod tidy
```

### ▶ Lancer l'application
```sh
go run main.go
```

L'application démarre un serveur HTTP sur **`http://localhost:8080`**.

### 🌐 Accéder à l'interface Web
Ouvrez votre navigateur et allez sur :
```
http://localhost:8080/monitor
```

## 📡 API Endpoints

### 🔹 CPU
- `GET /system/cpu` → Utilisation du CPU global
- `GET /system/cpu/cores` → Utilisation par cœur

### 🔹 Mémoire
- `GET /system/memory` → Stats mémoire

### 🔹 Disque
- `GET /system/disk` → Stats disque

### 🔹 Réseau
- `GET /system/network` → Stats réseau

### 🔹 Ports
- `GET /system/ports` → Ports en écoute

### 🔹 Informations Système
- `GET /system/info` → Nom de l'hôte, OS, uptime...

### 🔹 Alertes
- `GET /alerts` → Liste des alertes critiques

## 🛠 Technologies utilisées
- **Go** (Golang) 🦫
- **gopsutil** (Librairie pour récupérer les stats système)
- **net/http** (Gestion du serveur et des routes)
- **HTML, CSS, JavaScript** (Interface web dynamique)

## 📖 Contributions
Les contributions sont les bienvenues ! Vous pouvez ouvrir une issue ou faire une pull request sur le repo GitHub.

## 📜 Licence
Ce projet est sous licence MIT.

---
🎯 **Développé par MederickBernier avec passion ❤️ en Go !** 🚀
