# 🌈 Discoscreen

> Un outil Go pour Windows qui applique un effet de jeu de lumière (discothèque) dynamique à votre écran via l'API Magnification.

![Platform](https://img.shields.io/badge/Platform-Windows-0078D6?style=for-the-badge&logo=windows)

---

## ⚡ Fonctionnalités

- **Jeu de lumière plein écran** : Applique un filtre dynamique et coloré directement sur votre affichage.
- **API Windows Native** : Utilise `Magnification.dll` sans surcouche lourde.
- **Faible consommation de ressources** : Environ 2,3 Mo de mémoire et une faible utilisation du processeur, pour rester adapté aux PC moins puissants.
- **Arrêt rapide** : Appuyez sur la touche `Échap` pour arrêter l'effet.

## 🚀 Utilisation

Téléchargez [Discoscreen.exe depuis la release v1.0.0](https://github.com/l3nnyz/discoscreen/releases/tag/v1.0.0), puis lancez-le. Appuyez sur **Échap** pour arrêter l'effet et fermer le programme.

## 🛠️ Compiler le projet

Nécessite Windows et Go 1.26 ou une version ultérieure. Depuis le dossier du projet, lancez `build.bat`. Le fichier `discoscreen.exe` sera créé dans le même dossier.

## 🛠️ Structure du projet

```text
├── main.go      # Logique principale et boucle d'exécution
├── data.go      # Constantes et structures des effets lumineux
├── import.go    # Chargement dynamique des DLLs Windows
├── go.mod       # Gestion des dépendances du module
└── build.bat    # Script de compilation
```

> **Attention :** cet effet lumineux peut provoquer une gêne ou déclencher des crises chez les personnes sensibles aux lumières clignotantes.