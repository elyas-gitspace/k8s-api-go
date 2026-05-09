# Étape 1 — Build
FROM golang:1.21-alpine AS builder 
#docker crée un conteneur temporaire dans lequel il va faire tourner une machine avec Go installé

WORKDIR /app
#il se place dans le dossier app de ce conteneur temporaire

COPY go.mod .
COPY main.go .
#il copie le fichier go.mod et main.go dans le dossier app de ce conteneur temporaire

RUN go build -o api-go .
#il compile le code Go et crée l'exécutable api-go (en gros il traduit le langage Go en langage machine binaire pour que la machine puisse l'exécuter)

#La on va passer à la deuxième étape qui va être l'image finale. On a séparé le build et l'image finale pour optimiser le temps de build et l'espace disque utilisé. Ainsi on peut utiliser le conteneur temporaire pour le build et ensuite le conteneur final pour l'image finale qui sera plus légère et ne contiendra que le code compilé et les dépendances nécessaires, sans Go installé qui est lourd


# Étape 2 — Image finale
FROM alpine:latest
#on utilise une image alpine qui est une image linuxe (comme un pc avec un linux léger dedans) très légère et rapide
WORKDIR /app
#on se place dans le dossier app de ce conteneur final
COPY --from=builder /app/api-go .
#on copie l'exécutable api-go qu'on a créé dans le conteneur temporaire dans le dossier app de ce conteneur final
EXPOSE 8080
#on expose le port 8080 pour que le conteneur puisse être accessible depuis l'extérieur
CMD ["./api-go"]
#on lance l'exécutable api-go depuis ce conteneur final