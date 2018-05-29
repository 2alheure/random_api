# Alea Data Est

Ceci est un projet porté par Jordan Juventin (aka 2alheure) et Florian Rambur, étudiants à l'ECV Digital Paris. 
Il s'agit, principalement, d'une API fournissant des jeux de données générés aléatoirement. 

## Documentation

Retrouvez la documentation de l'API [ici](http://2dtension.fr/alea-data-est).  
Elle a été générée en utilisant [apidoc](http://apidocjs.com).


## TODO

- [x] type struct des models
- [x] fonctions d'hydratation des type struct
- [ ] fonctions de génération des model.type
- [ ] fonction de génération d'une ressource (en parcours en profondeur / largeur)
- [x] fonction helper d'output JSON
- [x] fonctions qui handle une requête HTTP
- [x] routes HTTP
  
## Evolutions ?

La liste des évolutions possibles est trop importante.  
Pour en avoir un aperçu, vous pouvez lire le cahier des charges.  
  
Néanmoins, pour le plus important, il y a un système d'utilisateur.  
Autrement dit, une authentification pour modifier la ressource propriétaire et un ratio de génération pour protéger d'un DDOS.  
  
En seconde position vient, indéniablement, une meilleure implémentation. Un meilleur algorithme, un code plus propre, plus encapsulé.  
  
Enfin, En troisième position, avant tout le reste : l'interface web. Parce que c'est quand même plus plaisant d'utiliser un service web par son interface graphique qu'en lignes de commande.