# Service Hôte ouvert

_DDDViteFait.pdf, p.76_

Lorsqu’on essaie d’intégrer deux sous-systèmes, on crée souvent une couche de traduction entre les deux. Cette couche agit comme un tampon entre le sous-système client et le sous-système externe avec lequel on essaie de s’intégrer. Cette couche peut être plus ou moins cohérente, selon la complexité des relations et la manière dont le sous-système externe a été conçu. Si ce dernier s’avère être utilisé non pas par un sous-système client, mais par plusieurs, on doit créer une couche de traduction pour chacun d’entre eux. Toutes ces couches vont répéter la même tâche de traduction, et contenir du code similaire.

Quand un sous-système doit être intégré avec beaucoup d’autres, fabriquer un traducteur sur mesure pour chacun peut enliser l’équipe. Il y a de plus en plus de choses à maintenir, et de plus en plus de choses dont il faut se soucier quand des modifications sont faites.

La solution est de voir le sous-système externe comme un fournisseur de services. Si nous pouvons l’enrober d’un ensemble de Services, alors tous les autres sous-systèmes accèderont à ces Services, et nous n’aurons pas besoin de couche de traduction. La difficulté est que chaque sous-système peut avoir besoin d’interagir de manière spécifique avec le sous-système externe, et créer un ensemble de Services cohérent peut s’avérer problématique.

Définissez un protocole qui donne accès à votre sous-système comme un ensemble de Services. Ouvrez le protocole pour que tous ceux qui doivent s’intégrer avec lui puissent l’utiliser. Améliorez et étendez le protocole pour gérer de nouvelles nécessités d’intégration, sauf quand une équipe a des besoins idiosyncratiques. Dans ce cas, utilisez un traducteur exceptionnel séparé pour élargir le protocole à cette situation particulière tout en gardant un protocole partagé simple et cohérent.
