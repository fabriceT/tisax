# Évaluation TISAX du site d'AGLA FORM

La sécurité d'un système d'information se résume dans l'objectif de la question **1.2.1** : 

> Le système de gestion de la sécurité de l’information (SMSI) est un mécanisme de contrôle utilisé par la direction de l’organisation pour s’assurer que la sécurité de l’information est le résultat d’une gestion durable plutôt que celle d’une simple coïncidence ou d’un effort individuel.

Ce document a pour but d'établir le niveau de maturité du système d'information du site d'AGLAFORM selon les critères de la labellisation TISAX dans sa version 5.1. 

A> Les version officielles de la grille d'évaluation étant en langue anglaise ou allemande, le document reprend les questions dans la version anglaise.

TISAX s'appuie sur les éléments suivants :

* ISO 27001 - Management de la sécurité de l'information 
* Des spécificités liées au monde le l'industrie automobile : la protection des prototypes
* le règlement général de protection des données

L'évaluation est réalisée sur le niveau de maturité du système d'information :

{.evaltext1}
1. **Effectué** - Un processus est suivi qui n’est pas ou pas suffisamment documenté (« processus informel »), il existe des preuves qu’il atteint son objectif.

{.evaltext2}
2. **Géré** - Le process est suivi et atteint ses objectifs. La documentation du processus et les preuves de la mise en œuvre du processus sont disponibles.

{.evaltext3}
3. **Établi** - Un processus standard intégré dans l’ensemble du système est suivi. Les dépendances à d’autres processus sont documentées et des interfaces appropriées sont créées. Il existe des preuves que le processus a été utilisé de manière durable et active sur une longue période.

Un liseret de couleur apparait dans la description de la question. Il varie selon le niveau d'évaluation

{.info .msgblock .center}
Dans cette évaluation, la note est de 0 (couleur noire) lorsqu'il n'y a pas de trace de processus.\
\
Certaines questions ne peuvent pas être évaluées car leur dépendance n'est pas satisfaite ou parce qu'il n'y pas de mesures se rapprochant de ce qui est demandé.

Des informations complémentaires sur les niveaux de maturité sont disponibles dans la [grille d'évaluation TISAX](https://portal.enx.com/isa5-en.xlsx). Une version plus détaillée et, complète, dans [AUTOMOTIVE SPICE v3.1 POCKET GUIDE](https://www.kuglermaag.com/fileadmin/05_CONTENT_PDF/literature_automotive-spice_pocketguide.pdf), p. 94 à p. 102.


Cette évaluation a été effectuée durant le mois de septembre 2022 par Fabrice THIROUX avec la participation de Marc LAFARGUE selon la grille d'évaluation TISAX disponible sur https://portal.enx.com/isa5-en.xlsx.

## Information Security
### 1) IS Policies and Organization
#### 1.1.1) To what extent are information security policies available?

{.evaltext0}
* <span>Objectif</span> : The organization needs at least one information security policy. This reflects the importance and significance of information security and is adapted to the organization. Additional policies may be appropriate depending on the size and structure of the organization.
* <span>Référence</span> : ISO 27001: A.5.1.1, A.5.1.2
* <span>Maturité</span> : 0 ❌

**Exigence** :

*  *Politiques de sécurité de l’information* : Il convient de définir un ensemble de politiques en matière de sécurité de l’information qui soit approuvé par la direction, diffusé et communiqué aux salariés et aux tiers concernés (A.5.1.1)
* *Revue des politiques de sécurité de l'information* : Pour garantir la constance de la pertinence, de l’adéquation et de l’efficacité des politiques liées à la sécurité de l’information, il convient de revoir ces politiques à intervalles programmés ou en cas de changements majeurs (A.5.1.2).

**Observations**

Pas de politique de sécurité des systèmes d'informations (PSSI). 

**Préconisations**

La mise en place d'une PSSI est nécessaire pour définir les exigences de sécurité de l'information du groupe, afficher sa volonté d'amélioration et du développement de la sécurité de son système de management de la sécurité de l'information (SMSI). La PSSI permet également d'assoir un ensemble de bonnes pratiques concernant l'hygiène de l'information pour l'ensemble des utilisateurs.

Le règles doivent être revues régulièrement ou si elles ne correspondent plus à l'état actuel du SMSI.

{.warning .msgblock}
La PSSI est la première brique vers la sécurisation du système d'information.




***
#### 1.2.1) To what extent is information security managed within the organization?

{.evaltext0}
* <span>Objectif</span> : Only if information security is part of the strategic goals of an organization, information security can be implemented in an organization in a sustainable manner. The information security management system (ISMS) is a control mechanism used by the organization’s management to ensure that information security is the result of sustainable management rather than that of mere coincidence and individual effort.
* <span>Référence</span> : ISO 27001: 4
* <span>Maturité</span> : 0 ❌


**Observations** :

L'organisation de l'informatique au sein du groupe F2J repose sur plusieurs entités :

* Des administrateurs ou ingénieur informatique sur site (M. LAFARGUE, M. TRANSLER, M. STRUGALA) travaillant chacun pour leur société respective (AGLA FORM, F2J REMAN CHAUMONT, F2J JAPY) sur des rôles de techniciens ou d'administrateurs systèmes. 
* M. COSSAIS, F2J DEVELOPMENT, en tant que responsable du système informatique pour la partie française du groupe. 

M. STRUGALA fait également office de support technique pour la partie française du groupe (AGLA FORM, F2J REMAN CHAUMONT, F2J JAPY et F2J STAMPING BESSINES) car M. LAFARGUE et M. TRANSLER occupent des fonctions hors informatique (Coordinateur HSE, logistique) et ne possèdent pas de solide formation en informatique.

M. COSSAIS participe à la modernisation de l'informatique du groupe et conduit les nombreux projets relatifs à cette modernisation. Il intervient également sur les différents sites français afin de réaliser des interventions du domaine d'un technicien.

On ne distingue pas d'organisation claire et précise au sein des personnes en charge de l'informatique. L'absence de SMSI ne permet pas de dégager des rôles précis ou des responsabilités au sein de l'organisation.

**Préconisations** :

La mise en place d'une politique de sécurité, appelée à évoluer relativement à l'état de l'art du moment, pour les systèmes d'informations (PSSI) et le maintient en conformité du système d'information vis-à-vis de cette politique nécessite la *désignation d'un responsable de la sécurité des systèmes d'information* (RSSI).

Le RSSI est le pivot de toutes actions ou choix touchant profondément à la sécurité du système d'information (réseaux, système, télécommunication, applications, sécurité physique...). Il doit également mettre en place un plan de continuité d'activité (PCA) et des moyens pour un fonctionnement en mode dégradé en cas de crise (rançongiciel, catastrophe naturelle…)

* [Memento cybersécurité pour le dirigeant d'entreprise](https://www.economie.gouv.fr/files/files/PDF/2017/bro-memento-cybersecurite-dirigeant_0.pdf)

#### 1.2.2) To what extent are information security responsibilities organized?

{.evaltext0}
* <span>Objectif</span> : A successful ISMS requires clear responsibilities within the organization.
* <span>Référence</span> : ISO 27001: A.6.1.1, A.6.1.2
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* *Fonctions et responsabilités liées à la sécurité de l’information* : Il convient de définir et d’attribuer toutes les responsabilités en matière de sécurité de l’information (A.6.1.1).
* *Séparation des tâches* : Il convient de séparer les tâches et les domaines de responsabilité incompatibles pour limiter les possibilités de modification ou de mauvais usage, non autorisé(e) ou involontaire, des actifs de l’organisation (A.6.1.2)

**Observations** :

L'organisation du système d'information repose, au niveau du groupe (cf [1.2.1](#toc_4)), sur un aspect fonctionnel au lieu de considérer l'aspect sécurité. Il en découle de nombreuses carences sur le plan de la sécurité de l'information car, contrairement à l'adage « la sécurité, c'est l'affaire de tous », la sécurité est ici l’affaire des personnes concernées, avec leur degré de sensibilisation et de moyens.

Ces carences ont de nombreuses répercussions sur l'évaluation TISAX, car l'évaluation de la partie sécurité de l'information repose sur des élements de la norme ISO-27001 qui :

* rationalise la sécurité
* impose la mise en place et le suivi de bonnes pratiques axées sur le plan de la sécurité (amélioration continue)
* permet de prendre en compte la sécurité à tous les niveaux.

**Préconisations** :

Une PSSI est le premier élément nécessaire pour débuter l'orientation de l'organisation. Elle sera l'affirmation claire de la politique de sécurité au niveau groupe dans le but d'entraîner le changement au sein d'AGLA FORM et des différentes sociétés du groupe F2J.

Les utilisateurs doivent être sensibilisés afin de comprendre leurs responsabilités et les contraintes liées.

Les responsabilités doivent être établies afin de prendre en charge leurs spécificités dans le système d'information.

Une cartographie du SI doit être réalisée afin de déterminer les points sensibles.

* [Lignes directrices de l'OCDE régissant la sécurité des systèmes et réseaux d'information : vers une culture de la sécurité](https://www.oecd.org/digital/ieconomy/15582260.pdf)
* [ANSSI: La cybersécurité des systèmes industriels](https://www.ssi.gouv.fr/guide/la-cybersecurite-des-systemes-industriels/)

#### 1.2.3) To what extent are information security requirements taken into account in projects?

{.evaltext0}
* <span>Objectif</span> : For project implementation, it is important to consider the information security requirements. This applies to projects within the organization regardless of their type. By appropriately establishing the information security process in the project management procedures of the organization, any overlooking of requirements is prevented.
* <span>Référence</span> : ISO 27001: A.6.1.5
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* *La sécurité de l'information dans la gestion de projet* : Il convient de traiter la sécurité de l’information dans la gestion de projet, quel que soit le type de projet concerné (A.6.1.5).

**Observations** :

Pas de politique de sécurité de l'information au niveau de la société ; Il est alors à considérer que la vision de la sécurité de l'information dans la gestion de projet équivaut à celle de son chef de projet.

**Préconisations** :

Mettre en place une vision globale de la sécurité avec une PSSI.

Tout projet doit faire l'objet d'une évaluation des risques (par exemple, un déménagement peut entrainer du passage et l'accès à des actifs protégés).

#### 1.2.4) To what extent are the responsibilities between external IT service providers and the own organization defined?

{.evaltext0}
* <span>Objectif</span> : It is important, that a common understanding of the division of responsibilities exists and that the implementation of all security requirements is ensured. Therefore, when using external IT service providers and IT services, the responsibilities regarding the implementation of information security measures are to be defined and verifiably documented.
* <span>Référence</span> : ISO 27017: CLD.6.3.1
* <span>Maturité</span> : 0 ❌


{.todo .msgblock}
Voir avec Serge s'il y a des services externalisés dont la sécurité de l'information n'est pas garantie (comme Mobility Work avec ses sauvegardes destructives et ses exports impossibles)

***
#### 1.3.1) To what extent are information assets identified and recorded?

{.evaltext1}
* <span>Objectif</span> : It is important for each organization to know the information constituting its essential assets (e.g. business secrets, critical business processes, know-how, patents). They are referred to as information assets. An inventory ensures that the organization obtains an overview of its information assets. Moreover, it is important to know the supporting assets (e.g. IT systems, services/IT services, employees) processing these information assets.
* <span>Référence</span> : ISO 27001: A.8.1.1, A.8.1.2
* <span>Maturité</span> : 1 ❌


**Exigences** :

* *Inventaire des actifs* : Il convient d’identifier les actifs associés à l’information et aux moyens de traitement de l’information et de dresser et tenir à jour un inventaire de ces actifs (A.8.1.1).
* *Propriété des actifs* : Il convient que les actifs figurant à l’inventaire aient un propriétaire. (A.8.1.2).

**Observations** :

Un suivi des actifs (ordinateurs, commutateurs) est en place. Ces actifs sont répertoriés et les utilisateurs identifiés dans un fichier Excel. La solution actuelle ne permet cependant pas de suivre la vie de l’élément (réparation, perte...).

**Préconisations** :

L'ISO 27001 définit comme actif tout ce qui peut avoir une "importance pour l'organisation". On retrouve donc dans l'ensemble de ces actifs :

* Matériel (ordinateurs, téléphones, imprimantes ou encore clefs USB);
* Logiciels (ensemble de logiciels utilisés);
* Information (bases de données, fichiers ou encore document papier);
* Infrastructure (électricité, climatisation), tout ce qui peut provoquer un problème sur le système d'information;
* Des personnes peuvent également y figurer car  disposant d'information ne figurant pas sous autre forme;
* Les services sous-traités peuvent également y trouver leur place (messagerie, service de nettoyage...).

Il faut identifier pour chaque actif son propriétaire. Il sera responsable de la protection de sa confidentialité, son intégrité et sa disponibilité.

Cet inventaire doit être régulièrement revu et mis à jour.

#### 1.3.2) To what extent are information assets classified and managed in terms of their protection needs?

{.evaltext0}
* <span>Objectif</span> : The objective of classifying information assets is the consistent determination of their protection needs. For this purpose, the value the information has for the organization is determined based on the protection goals of information security (confidentiality, integrity and availability) and classified according to a classification scheme. This enables the organization to implement adequate protective measures.
* <span>Référence</span> : ISO 27001: A.8.2.1, A.8.2.2, A.8.2.3, A.8.3.2
* <span>Maturité</span> : 0 ❌


**Exigences**:

* *Classification des informations* : Il convient de classer les informations en termes de valeur, d’exigences légales, de
sensibilité ou de leur caractère critique pour l’entreprise (A.8.2.1).
* *Marquage des informations* : Il convient d’élaborer et de mettre en oeuvre un ensemble approprié de procédures pour le marquage de l’information, conformément au plan de classification de l’information adopté par l’organisation (A.8.2.2).
* *Manipulation des actifs* : Il convient d’élaborer et de mettre en oeuvre des procédures de traitement des actifs, conformément au plan de classification de l’information adopté par l’organisation (A.8.2.3).
* *Mise au rebut des supports* : Il convient de procéder à une mise au rebut sécurisée des supports qui ne servent plus, en suivant des procédures formelles (A.8.3.2).

**Préconisations** :

Les informations doivent être classifiées par niveau de protections et marquées selon leur niveau. Des procédures de traitement de l'information, ainsi qu'une matrice des droits d'accès aux informations doivent être mis en place. 

Des procédures doivent être mises en place pour : 

* l'utilisation correcte des supports (accès restreint à une base de données, pas d'usage perso pour un PC…);
* le stockage des supports (coffre-fort, salle sécurisée, antivol…);
* la gestion de la vie des supports (retour, destruction, suppression des données…).

Des mesures doivent prises en compte pour restreindre l'accès aux actifs en fonction de leur niveau de confidentialité.


Niveau       | Description
:-----------:|--------------------------------------------------------------------------------------------------------------------------------------------------
Public       | Documents ou informations reconnus comme appartenant du domaine public
Interne      | Documents à usage interne et dont la divulgation aurait des impacts modéré
Confidentiel | Documents à destination d’un nombre de destinataires restreint et dont la divulgation aurait des impacts élevés.
Secret       | Documents à destination d’un nombre de destinataires restreint et nominativement identifiés et dont la divulgation aurait des impacts très élevés
Table: Les quatre niveaux de protection couramment proposés


* [Cybersécurité : classifier et protéger ses données](hhttps://www.afg.asso.fr/wp-content/uploads/2020/10/201014guideprodonnees-web.pdf)

#### 1.3.3) To what extent is it ensured that only evaluated and approved external IT services are used for processing the organization’s information assets?

{.evaltext0}
* <span>Objectif</span> : Particularly in the case of external IT services that can be used at relatively low cost or free of charge, there is an increased risk that procurement and commissioning will be carried out without appropriate consideration of the information security requirements and that security therefore is not ensured.
* <span>Référence</span> : ISO 27017: 14.1.1
* <span>Maturité</span> : 0 ❌


**Exigence** : 

* Veiller à ce que la sécurité de l’information fasse partie intégrante des systèmes d’information tout au long de leur cycle de vie. Cela inclut également des exigences pour les systèmes d’information fournissant des services sur les réseaux publics. (1.14.1)
  * *Analyse et spécification des exigences de sécurité de l’information* : Il convient que les exigences liées à la sécurité de l’information figurent dans les exigences des nouveaux systèmes d’information ou des changements apportés aux systèmes existants (A.14.1.1).
  * *Sécurisation des services d’application sur les réseaux publics* : Il convient de protéger l’information liée aux services d’application transmise sur les réseaux publics contre les activités frauduleuses, les différends contractuels, ainsi que la divulgation et la modification non autorisées (A.14.1.2)
  * *Protection des transactions liées aux services d’application* : Il convient de protéger l’information impliquée dans les transactions liées aux services d’application pour empêcher une transmission incomplète, des erreurs d’acheminement, la modification non autorisée, la divulgation non autorisée, la duplication non autorisée du message ou sa réémission. (A.14.1.3)

**Préconisations** :

L'intégration de nouveaux services, d'un nouvel applicatif ou d'améliorations au système d'information doit faire l'objet d'une évaluation des risques. Les exigences légales, réglementaires et contractuelles doivent être prises en compte.

Les services gratuits de partage de fichiers ou d'opération sur fichiers (conversion de formats, fusion PDF…) ne doivent pas être utilisés pour des documents internes à la société. Dans l'idéal, le service informatique doit fournir une liste des services approuvés, sous quelles conditions (par exemple, on ne partage pas sur Google Drive d'un compte personnel), ainsi que la correspondance avec les services qui ne doivent plus être utilisés.

Dans le cas d'utilisation de partages de fichiers sur des sites ne faisant pas partie du système d'information, les données non public téléversées sur ces sites doivent être chiffrées afin de préserver la confidentialité. Et, s'il y a utilisation de chiffrement symétrique, l'échange de la clef doit se faire par un autre support (un mail par exemple). Les tiers partageant des informations avec le groupe F2J via ces services doivent être informés du risque lié à leur utilisation : c'est à dire aucune confidentialité.

***
#### 1.4.1) To what extent are information security risks managed?

{.evaltext0}
* <span>Objectif</span> : Information security risk management aims at the timely detection, assessment and addressing of risks in order to achieve the protection goals of information security. It thus enables the organization to establish adequate measures for protecting its information assets under consideration of the associated prospects and risks. It is recommended to keep the information security risk management of an organization as simple as possible such as to enable its effective and efficient operation.
* <span>Référence</span> : ISO 27001: 6.1.2, 6.1.3
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* *Séparation des tâches* : Il convient de séparer les tâches et les domaines de responsabilité incompatibles
pour limiter les possibilités de modification ou de mauvais usage, non autorisé(e) ou involontaire, des actifs de l’organisation (A.6.1.2)
* *Relations avec les autorités* : Il convient d’entretenir des relations appropriées avec les autorités compétentes (A.6.1.3).

**Préconisations** :

L'identification des tâches incompatibles et attributions des responsabilités en conséquence doivent permettre de mettre en places des mesures pour cloisonner les domaines de responsabilités incompatibles et limiter les possibilités de modification ou des mauvais usages (non autorisé ou involontaire) des actifs de l'organisme.

* L'organisation de la gestion des risques du système d'information doit être clairement établie;
* Un inventaire des actifs doit être réalisé (au moins ceux portant sur des informations confidentielles);
* Des fiches de poste doivent être établies afin de recenser les tâches incompatibles.

Le site [Cybermalveillance](https://www.cybermalveillance.gouv.fr) liste les démarches à suivre en cas d'incident, la partie [En cas d'incident](https://www.ssi.gouv.fr/en-cas-dincident/) de l'ANSSI également.

***
#### 1.5.1) To what extent is compliance with information security ensured in procedures and processes?

{.evaltext0}
* <span>Objectif</span> : It is not sufficient to define information security requirements and to prepare and publish policies. It is important to regularly review their effectiveness.
* <span>Référence</span> : ISO 27001: A.18.2.2, A.18.2.3
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* *Conformité avec les politiques et les normes de sécurité* : Il convient que les responsables revoient régulièrement la conformité du
traitement de l’information et des procédures dont ils sont chargés au regard des politiques, des normes de sécurité applicables et autres exigences de sécurité (A.18.2.2)
* *Examen de la conformité technique* : Il convient que les systèmes d’information soient régulièrement revus pour vérifier leur conformité avec les politiques et les normes de sécurité de l’information de l’organisation (1.18.2.3)

**Observations** :

Pas de politique de sécurité existante.

**Préconisations** :

Définir les responsabilités dans [1.2.2](#toc_5).

Les responsables doivent régulièrement vérifier la conformité du traitement de l’information et des procédures dont ils sont chargés au regard des politiques, des normes de sécurité applicables et autres exigences de sécurité. 

Un audit régulier permet de s'assurer de la conformité du système d'information.

#### 1.5.2) To what extent is the ISMS reviewed by an independent authority?

{.evaltext0}
* <span>Objectif</span> : As an essential control mechanism, assessing the effectiveness of the ISMS from merely an internal point of view is insufficient. Additionally, an independent and therefore objective assessment shall be obtained at regular intervals and in case of significant changes.
* <span>Référence</span> : ISO 27001: A.18.2.1
* <span>Maturité</span> : 0 ❌


**Exigence** : 

* *Revue indépendante de la sécurité de l’information* : Il convient de procéder à des revues régulières et indépendantes de l’approche
retenue par l’organisation pour gérer et mettre en oeuvre la sécurité de l’information (à savoir le suivi des objectifs, les mesures, les politiques, les procédures et les processus relatifs à la sécurité de l’information) à intervalles définis ou lorsque des changements importants sont intervenus. (A.18.2.1)

**Préconisations** :

Les systèmes d’information doivent être examinés régulièrement, de manière indépendante, quant à leur conformité avec les politiques et les normes de sécurité de l’information de l’organisme.

***
#### 1.6.1) To what extent are information security events processed?

{.evaltext0}
* <span>Objectif</span> : Organized processing of information security events aims at limiting potential damage and preventing recurrence.
* <span>Référence</span> : ISO 27001: A.16.1.1, A16.1.2
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* *Responsabilités et procédures* : Il convient d’établir des responsabilités et des procédures permettant de garantir une réponse rapide, efficace et pertinente en cas d’incident lié à la sécurité de l’information. (A.16.1.1)
* *Signalement des événements liés à la sécurité de l’information* : Il convient de signaler, dans les meilleurs délais, les événements liés à la sécurité de l’information, par les voies hiérarchiques appropriées. (A.16.1.2)

**Observations** :

Pas de gestion des incidents, ni de procédures pour y répondre. 

Le signalement des incidents, liés à la sécurité informatique ou non, est laissé au bon vouloir de l'utilisateur. Les utilisateurs doivent être informés de leur obligation à déclarer les incidents.

**Préconisations** :

Une obligation de signaler un incident doit être effective et connue de tous.

Une procédure de signalement d'incident doit être mise en place et diffusée (par exemple, création d'une adresse mail pour le signalement).

Les utilisateurs doivent être formés pour réagir de manière correcte en cas d'incident (par exemple, déconnexion du réseau en cas de rançongiciel) afin de limiter les dégats.

***
### 2) Human Resources
#### 2.1.1) To what extent is the qualification of employees for sensitive work fields ensured?

{.evaltext0}
* <span>Objectif</span> : Competent, reliable and trustworthy employees are a key to information security within the organization Therefore, it is important to check the qualifications of potential employees (e.g. applicants) to an appropriate extent.
* <span>Référence</span> : ISO 27001: A.7.1.1
* <span>Maturité</span> : 0 ❌

**Exigence** :

* *Sélection des candidats* : Il convient que des vérifications des informations concernant tous les candidats à l’embauche soient réalisées conformément aux lois, aux règlements et à l’éthique, et il convient qu’elles soient proportionnelles aux exigences métier, à la classification des informations accessibles et aux risques identifiés. (A.7.1.1).

#### 2.1.2) To what extent is all staff contractually bound to comply with information security policies?

{.evaltext1}
* <span>Objectif</span> : Organizations are subject to legislation, regulations and internal policies. Already when hiring staff, it must be ensured that employees commit to compliance with the policies and are aware of the consequences of misconduct.
* <span>Référence</span> : ISO 27001: A.7.1.2, A.7.3.1
* <span>Maturité</span> : 1 ❌


**Exigences** :

* *Termes et conditions d’embauche* : Il convient que les accords contractuels conclus avec les salariés et les contractants déterminent leurs responsabilités et celles de l’organisation en matière de sécurité de l’information (A.7.1.2)
* *Achèvement ou modification des responsabilités associées au contrat de travail* : Il convient de définir les responsabilités et les missions liées à la sécurité de l’information qui restent valables à l’issue de la rupture, du terme ou de la modification du contrat de travail, d’en informer le salarié ou le contractant et de veiller à leur application (A.7.3.1)

**Observations** :
Les personnes ayant un accès aux données confidentielles dans le cadre de leurs fonctions doivent signer une lettre de confidentialité.

{.info .msgblock}
Malgré cela, intérimaire, et étant titulaire d'un compte utilisateur ayant les pouvoirs d'administration et ayant l'accès à l'intégralité des données informatiques stockées sur les serveurs du site, je n'ai jamais signé de lettre de confidentialité.

**Préconisations** :

* Définir les responsabilités dans [1.2.2](#toc_5).

#### 2.1.3) To what extent is staff made aware of and trained with respect to the risks arising from the handling of information?

{.evaltext1}
* <span>Objectif</span> : If the requirements and risks of information security are not known to the employees, there is a risk of misconduct resulting in damage to the organization. Therefore, it is important that information security is internalized and practiced as a natural part of their work.
* <span>Référence</span> : ISO 27002: A.7.2.1, A.7.2.2
* <span>Maturité</span> : 1 ❌


**Exigences** : 
 
* *Responsabilités de la direction* : Il convient que la direction demande à tous les salariés et contractants
d’appliquer les règles de sécurité conformément aux politiques et aux procédures en vigueur dans l’organisation (A.7.2.1).
* *Sensibilisation, apprentissage et formation à la sécurité de l’information* : Il convient que l’ensemble des salariés de l’organisation et, le cas échéant, les contractants suivent un apprentissage et des formations de sensibilisation adaptés et qu’ils reçoivent régulièrement les mises à jour des politiques et procédures de l’organisation s’appliquant à leurs fonctions (A.7.2.2).

**Observations** :

Des informations sur les risques critiques du moment sont dispensées par messagerie électronique (par exemple, vague d’hameçonnage). 

Pas de formation sur les risques liés à la sécurité informatique, pas de politique sur la sécurité de l'information.

**Préconisations** : 

* Réaliser une formation pour l'ensemble du personnel utilisant un ordinateur, par exemple le [MOOC de l'ANSSI](https://secnumacademie.gouv.fr/).
* Des formations complémentaires devraient être dispensées pour les personnes en charge de l'ISMS afin de maintenir leurs compétences et système d'information au niveau exigé par le contexte du moment (cybercriminalité, réglementation...).

[La formation du personnel peut stopper 59% des incidents de sécurité](https://advisera.com/27001academy/blog/2015/02/16/change-thinking-can-stop-59-security-incidents/).

#### 2.1.4) To what extent is teleworking regulated?

{.evaltext0}
* <span>Objectif</span> : Working outside the specifically defined security zones (teleworking) creates particular risks requiring corresponding protective measures.
* <span>Référence</span> : ISO 27001: A.6.2
* <span>Maturité</span> : 0 ❌


**Exigence**

* *Télétravail* : Il convient de mettre en œuvre une politique et des mesures de sécurité complémentaires pour protéger les informations consultées, traitées ou stockées sur des sites de télétravail (A.6.2.2)

**Observation**

Pas de mesure.

**Préconisation**

Des mesures doivent être prise pour :

* limiter l'accès aux appareils appartenant à la société
* Sensibiliser les utilisateurs aux bonnes conduites à tenir lors du télétravail (pas de partage de l'appareil avec des tiers, séparation des usages)

***
### 3) Physical Security and Business Continuity
#### 3.1.1) To what extent are security zones managed to protect information assets?

{.evaltext0}
* <span>Objectif</span> : Security zones provide physical protection of information assets. The more sensitive the information assets to be processed are the more protective measures are required.
* <span>Référence</span> : ISO 27001: A.11.1
* <span>Maturité</span> : 0 ❌

**Exigences** :

* *Zones sécurisées* :Empêcher tout accès physique non autorisé, tout dommage ou intrusion portant sur l’information et les moyens de traitement de l’information de l’organisme (A.11.1).
  * *Périmètre de sécurité physique* : Il convient de définir des périmètres de sécurité servant à protéger les zones contenant l’information sensible ou critique et les moyens de traitement de l’information (A.11.1.1)
  * *Contrôle d’accès physique* : Il convient de protéger les zones sécurisées par des contrôles adéquats à l’entrée pour s’assurer que seul le personnel autorisé est admis (A.11.1.2)
  * *Sécurisation des bureaux, des salles et des équipements* : Il convient de concevoir et d’appliquer des mesures de sécurité physique aux bureaux, aux salles et aux équipements. (A.11.1.3)
  * *Protection contre les menaces extérieures et environnementales* : Il convient de concevoir et d’appliquer des mesures de protection physique contre les désastres naturels, les attaques malveillantes ou les accidents (A.11.1.4).
  * *Travail dans les zones sécurisées*: Il convient de concevoir et d’appliquer des procédures pour le travail en zone
sécurisée (A.11.1.5)
  * *Zones de livraison et de chargement* : Il convient de contrôler les points d’accès tels que les zones de livraison et de
chargement et les autres points par lesquels des personnes non autorisées peuvent pénétrer dans les locaux et, si possible, de les isoler des moyens de traitement de l’information, de façon à éviter les accès non autorisés. (A.11.1.6)

**Observations** :

Pas de mesure de sécurité concernant l'accès à la salle serveur.

{.todo .msgblock}
Voir avec Serge si les bureaux ferment à clefs ou si les visiteurs sont accompagnés.

**Préconisations**

Mettre en place des mesures de sécurisation des lieux contenant de l'information :

* Les livraisons ne doivent être réalisées dans les lieux où se trouvent des actifs sensibles (ex: pas de livreurs dans les bureaux)
* les endroits où se trouvent des actifs sensibles doivent être protégés afin que seul le personnel autorisé puisse y avoir accès (fermeture à clef, film occultant...)
* Les actifs doivent être protégés des dommages environnementaux (eau, incendie...), ainsi que des dommages humains (intrusion, café...)

Un contrôle d'accès doit être mis en place pour prévenir toute présence de personnes non autorisées dans les espaces pouvant contenir des actifs. 

Dans le cas d'un contrôle d'accès par carte, son usage doit être supervisé afin de contrôler son efficience (par exemple, sas d'accès ouvert pour faciliter le passage durant les heures de travail).

#### 3.1.2) To what extent is information security ensured in exceptional situations?

{.evaltext1}
* <span>Objectif</span> : Exceptional situations (e.g. natural disasters, physical attacks, cyber attacks, exceptional social situations, incidents or infrastructure failures of significant impact) present a great challenge to the organization. Good preparation helps to ensure that information security risks are adequately considered even in exceptional situations.
* <span>Référence</span> : ISO 27001: A.12.3, A.17.1, A.17.2
* <span>Maturité</span> : 1 ❌


**Exigences** :

* *Sauvegarde* : Se protéger de la perte de données (A.12.3).
  * *Sauvegarde des informations* : Il convient de réaliser des copies de sauvegarde de l’information, des logiciels et des images systèmes, et de les tester régulièrement conformément à une politique de sauvegarde convenue. (A.12.3.1)
* *Continuité de la sécurité de l’information* : Il convient que la continuité de la sécurité de l’information fasse partie intégrante des systèmes de gestion de la continuité de l’activité. (A.17.1)
  * *Organisation de la continuité de la sécurité de l’information* : Il convient que l’organisation détermine ses exigences en matière de sécurité de
l’information et de continuité du management de la sécurité de l’information dans des situations défavorables, comme lors d’une crise ou d’un sinistre (A.17.1.1)
  * *Mise en œuvre de la continuité de la sécurité de l’information* : Il convient que l’organisation établisse, documente, mette en œuvre et maintienne à jour des processus, des procédures et des mesures permettant de garantir le niveau requis de continuité de la sécurité de l’information au cours d’une situation défavorable (A.17.1.2)
  * *Vérifier, revoir et évaluer la continuité de la sécurité de l’information* : Il convient que l’organisation vérifie à intervalles réguliers les mesures de
continuité de la sécurité de l’information déterminées et mises en œuvre, afin de s’assurer qu’elles restent valables et efficaces dans des situations défavorables (A.17.1.3)
* *Redondances* : Garantir la disponibilité des moyens de traitement de l’information (A.17.2)
  * *Disponibilité des moyens de traitement de l’information* : Il convient de mettre en œuvre des moyens de traitement de l’information avec suffisamment de redondances pour répondre aux exigences de disponibilité (A.17.2.1)

**Observations** :

Les sauvegardes sont stockées sur un NAS Synology. Le NAS n'a pas été mis à jour depuis 2018 et présente de nombreuses failles de sécurité.

Une copie de la sauvegarde est réalisée manuellement sur un disque externe. L'opération manuelle ne permet pas de garantir que la copie soit identique à l'original, ni qu'elle soit à jour. Il n'est donc pas assuré que l'on puisse restaurer à partir de ce support

**Préconisations** :

Un inventaire des ressources à sauvegarder doit être réalisé afin que l'ensemble des informations importantes puissent être sauvegardées (ex. pas de données importantes stockées dans une clefs USB et oubliées du plan de sauvegarde).

Les sauvegardes doivent être testées régulièrement. Elles doivent être mises à l'abri dans un local sécurisé et être protégées contre les risques d'incendie ou des dégâts de eaux. Dans le cas où les sauvegardes permettent de reconstituer l'environnement de production, une copie doit être sauvegardée en dehors du site de production (par exemple, tâche de synchronisation automatique NAS -> disque externe stocké hors site, comme à F2J JAPY).

Une stratégie de sauvegarde doit être mise en place afin de répondre aux durées légales et aux exigences internes.

Un plan de continuité de l'activité ([PCA](https://www.economie.gouv.fr/files/hfds-guide-pca-plan-continuite-activite-_sgdsn.pdf)) doit être prévu. La continuité de la sécurité de l'information doit être établie, documentée à l'aide de procédures et de processus tenus à jour. 

Une structure de gestion de crise être mise en place afin que le niveau requis de sécurité de l'information puisse continuer à être fourni. Les procédures doivent être revues, testées et disponibles même en cas d'indisponibilité de l'infrastructure informatique. 

Les moyens de traitement de l’information doivent être mis en œuvre avec suffisamment de redondances pour répondre aux exigences de disponibilité ; La cartographie du SI doit permettre de détecter les points de défaillance uniques.

#### 3.1.3) To what extent is the handling of supporting assets managed?

{.evaltext0}
* <span>Objectif</span> : During their lifecycle (e.g. usage, disposal), supporting assets are subject to risks such as loss, theft or unauthorized viewing.
* <span>Référence</span> : ISO 27001:  A.8.1.3 and A.8.1.4
* <span>Maturité</span> : 0 ❌


**Exigences** :

* *Utilisation correcte des actifs* : Il convient d’identifier, de documenter et de mettre en œuvre des règles d’utilisation correcte de l’information, des actifs associés à l’information et des moyens de traitement de l’information. (A.8.1.3)
* *Restitution des actifs* : Il convient que tous les salariés et utilisateurs tiers restituent la totalité des actifs de l’organisation qu’ils ont en leur possession au terme de la période d’emploi, du contrat ou de l’accord (A.8.1.4)

**Observations**

Il n'existe pas de procédures du bon usage des actifs.

**Préconisations**

Les règles d’utilisation correcte de l’information, des actifs associés à l’information et des moyens de traitement de l’information doivent être identifiées, documentées et mises en œuvre.

Les salariés et utilisateurs tiers doivent restituer la totalité des actifs de l'organisme qu'ils ont au terme de la période d’emploi, du contrat ou de l’accord. Des documents doivent être mis en place pour suivre la gestion des actifs de type ordinateur ou téléphone (PV de réception, restitution, perte, réparation...)

#### 3.1.4) To what extent is the handling of mobile IT devices and mobile data storage devices managed?

{.evaltext0}
* <span>Objectif</span> : Mobile IT devices (e.g. notebooks, tablets, smartphones) and mobile data storage devices (e.g. SD cards, hard drives) are generally used not only on the premises of an organization, but also in mobile applications. This presents an increased risk with respect to e.g. loss or theft.
* <span>Référence</span> : ISO 27001: A.6.2, A.8.3
* <span>Maturité</span> : 0 ❌


**Exigences** :

* *Appareils mobiles et télétravail* : Assurer la sécurité du télétravail et l'utilisation d’appareils mobiles (A.6.2)
  * *Politique en matière d’appareils mobiles* : Il convient d’adopter une politique et des mesures de sécurité complémentaires
pour gérer les risques découlant de l’utilisation des appareils mobiles (A.6.2.1).
* *Manipulation des supports* : Empêcher la divulgation, la modification, le retrait ou la destruction non autorisé de l’information de l’organisme stockée sur des supports (A.8.3).
  * *Gestion des supports amovibles* : Il convient de mettre en œuvre des procédures de gestion des supports amovibles conformément au plan de classification adopté par l’organisation (A.8.3.1)
  * *Mise au rebut des supports* : Il convient de procéder à une mise au rebut sécurisée des supports qui ne servent
plus, en suivant des procédures formelle (A.8.3.2).
  * *Transfert physique des supports* : Il convient de protéger les supports contenant de l’information contre les accès non autorisés, l’utilisation frauduleuse ou l’altération lors du transport.

**Observations**

* Pas de MDM pour la gestion des périphériques mobiles (Téléphones, ordinateurs portables).
* Pas de chiffrement de la mémoire de stockage (Bitlocker...)

La perte ou le vol d'un téléphone mobile, d'un PC portable ou stockage amovible, peut entraîner une divulgation d'information sans possibilité de contenir la fuite.

**Préconisations**

a) Gestion des appareils mobiles :

Il est nécessaire de limiter les risques liés à la perte ou le vol d'un matériel mobile en chiffrant le contenu des disques. 

L'utilisation d'un MDM (Mobile Device Management) permet :

* de limiter les risques à l'usage en interdisant l'installation d'applications puisant dans les données de l'utilisateur
* supprimer le contenu en cas de perte.
* définir des politiques d'usage.

b) Gestion de l'information :

* Définir des procédures d'archivage, destruction des supports contenant de l'information confidentielle. 
* Protéger ces supports contre une lecture ou modification non autorisée.

Il est important d'avoir une classification de l'information afin de garantir que les informations confidentielles soient protégées. Il est également important que les utilisateurs soient sensibilisés sur le niveau de confidentialité des informations et actifs qu'ils traitent, ainsi que sur les procédures à suivre pour garantir leur confidentialité, intégrité, disponibilité et non-répudiation.

***
### 4) Identity and Access Management
#### 4.1.1) To what extent is the use of identification means managed?

{.evaltext1}
* <span>Objectif</span> : To check the authorization for both physical access and electronic access, means of identification such as keys, visual IDs or cryptographic tokens are often used. The security features are only reliable if the use of such identification means is handled adequately.
* <span>Référence</span> : ISO 27001: A.9.2.6
* <span>Maturité</span> : 1 ❌


**Exigences**

* *Suppression ou adaptation des droits d’accès* : Il convient que les droits d’accès de l’ensemble des salariés et utilisateurs tiers à l’information et aux moyens de traitement de l’information soient supprimés à la fin de leur période d’emploi, ou adaptés en cas de modification du contrat ou de l’accord (A.9.2.6).

**Observations**

On trouve des comptes d'utilisateur qui n'ont pas été connectés depuis plus de 120 jours. Dans ces comptes ayant un accès aux fichiers :

* utilisateur Test, dernière connexion 2015
* Freddy CORDIER, dernière connexion 2015
* depan01 (2010)
* Prince (2011)
* Stagiaire (mai 2022)
* Didier BASIN (2020), avec accès methodes, CAO, Production
* Pascal DELCOURT (2020)
* Sebastien GUCHE (2021)
* Leader ZF (2021)
* Verax Formage (2019)
* Giuliano Dallaricca (2019), avec des accès Projet, qualité, R&D, Qualité management
* Transfert (2020)
* transfert.labo (2021)
* HK 1329 (mai 2022)
* visio accueil (2021)
* tbi administration (2021)

Les comptes ne sont pas protégés contre les attaques informatiques : les mots de passe sont "d'époque" (complexité inconnue) et il n'y a pas de bloquage du compte en cas d'attaque par force brute.

On trouve beaucoup de comptes utilisateurs désactivés dans l'Active Directory (tech-form.local/Postes Fixes/Utilisateurs/Partis), peu ont cependant été désactivés ces 6 dernières années (Jean-Luc MAUCONDUIT, Rolland TINCHON).

Année | Nombre
:----:|:------:
2010  | 8
2011  | 4
2012  | 6
2013  | 7
2015  | 1
2016  | 8
2019  | 1
2022  | 1 
Table: Année de dernière connnexion pour les personnes désactivées.

On trouve également de nombreux comptes ordinateurs dans les comptes non désactivés. Ces comptes sont des vecteurs d'attaque car la relation entre le domaine et l'ordinateur n'a pas été coupé.

La présence de ces comptes ordinateurs implique qu'il n'y a pas de prise en compte du retrait des actifs (voir [1.3.2](#toc_9))

{.todo .msgblock}
Voir avec le RH pour les départs

**Préconisation**

L'arrivée, le départ ou le changement de position des utilisateurs doit être formalisé afin de prendre l'ensemble des entrées au système d'information accordés à l'utilisateur (accès, compte, matériel...).

#### 4.1.2) To what extent is the user access to network services, IT systems and IT applications secured?

{.evaltext1}
* <span>Objectif</span> : Only securely identified (authenticated) users are to gain access to IT systems. For this purpose, the identity of a user is securely determined by suitable procedures.
* <span>Référence</span> : IS0 27001: A.9.1, A.9.4.2
* <span>Maturité</span> : 1 ❌


**Exigence**

* *Exigences métier en matière de contrôle d’accès* : Limiter l'accès à l'information et aux moyens de traitement de l'information (A.9.1).
  * *Politique de contrôle d'accès* : Il convient d’établir, de documenter et de revoir une politique du contrôle d’accès sur la base des exigences métier et de sécurité de l’information. (A.9.1.1)
  * *Accès aux réseaux et aux services en réseau* : Il convient que les utilisateurs aient uniquement accès au réseau et aux services en réseau pour lesquels ils ont spécifiquement reçu une autorisation. (A.9.1.2)
* *Sécuriser les procédures de connexion* : Lorsque la politique de contrôle d’accès l’exige, il convient que l’accès aux systèmes et aux applications soit contrôlé par une procédure de connexion sécurisée (A.9.4.2)

**Observations**

Les utilisateurs sont authentifiés sur l'AD. La structure des groupes utilisateurs permet une granularité fine sur les droits d'accès. 
Peu de comptes génériques, les opérations réalisées dans le système d'information sont donc majoritairement nominatives.

L'absence de matrices des droits ne permet pas de confirmer que les accès sont conformes à ce qui est attendu.

Le partage `public` dont le contenu est en lecture/écriture pour les utilisateurs du système d'information présente plusieurs risques :

* le manque de politique d'usage peut entraîner un usage comme un espace de travail ou stockage.
* L'accès en lecture à l'ensemble des utilisateurs peut entraîner des fuites d'information.
* l'accès en écriture à l'ensemble des utilisateurs peut entraîner des pertes d'information, volontaires ou non.

Il n'y a pas de politique de verrouillage de compte en place. Il est possible pour un attaquant de tester un nombre 'infini' de mot de passe.

**Préconisations** :

La présence des fichiers dans le partage `public` devraient être limitées dans le temps et le rôle du partage bien établi.

Le compte administrateur est utilisé par plusieurs personnes. Prévoir un compte administrateur pour chaque personne devant réaliser des tâches d'administration.

Mettre un verrouillage de compte au bout de 3 ou 5 tentatives échouées.

Il est nécessaire de sécuriser l'accès au ressources stockées sur le cloud (messagerie, Drive...) pour les VIP, en mettant en place l'authentification forte (authentification multi-facteurs).

#### 4.1.3) To what extent are user accounts and login information securely managed and applied?

{.evaltext1}
* <span>Objectif</span> : Access to information and IT systems is provided via validated user accounts assigned to a person. It is important to protect login information and to ensure the traceability of transactions and accesses.
* <span>Référence</span> : IS0 27001: A.9.2.1, A.9.2.2, A.9.2.4, A.9.3.1, A.9.4.3
* <span>Maturité</span> : 1 ❌


**Exigences**

* *Enregistrement et désinscription des utilisateurs* : Il convient de mettre en œuvre une procédure formelle d’enregistrement et de
désinscription des utilisateurs destinée à permettre l’attribution de droits d’accès (A.9.2.1)
* *Maîtrise de la gestion des accès utilisateur* : Il convient de mettre en œuvre un processus formel de maîtrise de la gestion des
accès utilisateur pour attribuer ou révoquer des droits d’accès à tous les types d’utilisateurs de tous les systèmes et de tous les services d’information (A.9.2.2).
* *Gestion des informations secrètes d’authentification des utilisateurs*: Il convient que l’attribution des informations secrètes d’authentification soit réalisée dans le cadre d’un processus de gestion formel (A.9.2.4).
* *Utilisation d’informations secrètes d’authentification*: Il convient d’exiger des utilisateurs des informations secrètes d’authentification qu’ils appliquent les pratiques de l’organisation en la matière (A.9.3.1).
* *Système de gestion des mots de passe* : Il convient que les systèmes qui gèrent les mots de passe soient interactifs et fournissent des mots de passe de qualité (A.9.4.3).

**Observation**

L'attribution des accès aux partages est réalisée manuellement, ce type de traitement est enclin aux erreurs et aux oublis.

**Préconisations**

Une matrice des droits d'accès doit être réalisée afin de pouvoir déterminer automatiquement les droits d'accès aux partages en les liants à des rôles, service, positions... L'arborescence dans l'Active Directory doit permettre de définir ou révoquer automatiquement ces droits d'accès. Il ne restera alors qu'à gérer les exceptions à la règle.

Ces exceptions seront visibles dans la matrices des droits d'accès et traitées en conséquence lors de changements dans l'organisation.

Les utilisateurs doivent être sensibilisés au traitement d'information secrète :

* Pas de partages des identifiants
* Changement de mot de passe en cas de soupçon
* Ne pas stocker les identifiants dans des endroits accessibles (papier, fichier excel...)
* Ne pas réutiliser les mêmes informations secrètes sur différents systèmes d'authentification
* Savoir créer des mots de passe forts

***
#### 4.2.1) To what extent are access rights assigned and managed?

{.evaltext0}
* <span>Objectif</span> : The management of access rights ensures that only authorized users have access to information and IT applications. For this purpose, access rights are assigned to user accounts.
* <span>Référence</span> : ISO 27001: A.9.2.3, A.9.2.5, A.9.4.1
* <span>Maturité</span> : 0 ❌

**Exigence** 

* *Gestion des privilèges d’accès* : Il convient de restreindre et de contrôler l’attribution et l’utilisation des
privilèges d’accès (A.9.2.3).
* *Revue des droits d'accès utilisateurs* : Il convient que les propriétaires d’actifs revoient les droits d’accès des utilisateurs
à intervalles réguliers (A.9.2.5). 
* *Restriction d'accès à l'information* : Il convient de restreindre l’accès à l’information et aux fonctions d’application système conformément à la politique de contrôle d’accès.

**Préconisation**

La matrice des droit d'accès va servir de référence pour valider que les droits accordés sont conformes

Les responsable d'actifs doivent également s'assurer que l'usage de l'actif est conforme. Un nouvel utilisateur de l'actif doit être informé des usages et responsabilités liés à l'accès à l'actif.

***
### 5) IT Security / Cyber Security
#### 5.1.1) To what extent is the use of cryptographic procedures managed?

{.evaltext0}
* <span>Objectif</span> : When using cryptographic procedures, it is important to consider risks in the field of availability (lost key material) as well as risks due to incorrectly applied procedures in the fields of integrity and confidentiality (poor algorithms/protocols or insufficient key strengths).
* <span>Référence</span> : ISO 27001: A.10.1
* <span>Maturité</span> : 0 ❌


**Exigences** :

* *Mesures cryptographiques*: Garantir l’utilisation correcte et efficace de la cryptographie en vue de protéger la confidentialité, l’authenticité ou l’intégrité de l’information (A.10.1). 
  * *Politique d’utilisation des mesures cryptographiques* : Il convient d’élaborer et de mettre en œuvre une politique d’utilisation de mesures cryptographiques en vue de protéger l’information (A.10.1.1).
  * *Gestion des clés* : Il convient d’élaborer et de mettre en œuvre tout au long de leur cycle de vie une politique sur l’utilisation, la protection et la durée de vie des clés cryptographiques (A.10.1.2).

**Observation**

Pas de prise en compte de la cryptographie dans les échanges ou le stockage.

**Préconisation**

Les clefs cryptographiques, certificats, doivent avoir une durée de vie conforme à la politique prévue à cet effet.

Dans l'idéal, toute information sensible stockée sur un support mobile quittant ou pouvant quitter le lieu d'établissement devrait être chiffrée, soit en utilisant un support fournissant ce service (clef USB sécurisée), soit à l'aide d'un système de fichiers chiffré (Bitlocker), soit en utilisant une solution de chiffrement logiciel (conteneur, pgp...). Une politique d'utilisation doit être mise en place sur ce sujet.

L'utilisation d'un coffre-fort numérique (Keepass...) est recommandée pour stocker des clefs de déchiffrement et autres mot de passe. Il permet également de suivre leur durée de vie

* [Produits et services qualifiés par l'ANSSI](https://www.ssi.gouv.fr/uploads/liste-produits-et-services-qualifies.pdf)
* [KeepassXC](https://keepassxc.org/)
* [Veracrypt](https://veracrypt.fr/en/Home.html)

#### 5.1.2) To what extent is information protected during transfer?

{.evaltext0}
* <span>Objectif</span> : When being transferred via public or private networks, information can in some circumstances be read or manipulated by unauthorized third parties. Therefore, requirements regarding the protection needs of the information must be determined and implemented by taking suitable measures during such transfer.
* <span>Référence</span> : ISO 27001: A.13.2.1, A.13.2.3
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* *Politiques et procédures de transfert de l’information* : Il convient de mettre en place des politiques, des procédures et des mesures de transfert formelles pour protéger les transferts d’information transitant par tous types d’équipements de communication (A.13.2.1)
* *Messagerie électronique* : Il convient de protéger de manière appropriée l’information transitant par la messagerie électronique (A.13.2.3).

**Observation**

Pas de politique de chiffrement pour les données.

Pas de règles concernant la messagerie électronique

Il est impossible dans l'état actuel impossible de protéger la messagerie contre perte d'information. Le niveau de licence de Google Workspace n'offre pas la possibilité de définir une politique de rétention d'information, rendant de fait possible la suppression définitive (intentionnelle ou non) de messages électroniques.

**Préconisation**

La PSSI devrait prendre en compte l'utilisation de chiffrement sur les données partagées avec des tiers (cf. [1.3.3](#toc_10)).

L'utilisation de service web demandant l'échange d'information (mot de passe, document,...) doit se faire à l'aide d'un protocole chiffré (HTTPS)

La messagerie devrait être protégée à l'aide d'une authentification forte pour empêcher une usurpation. L'utilisation des services de messageries instantanée ou réseau sociaux doit être évaluée (cf. [1.3.3](#toc_10)).

Les services en ligne doivent pouvoir permettre la mise en place une politique de rétention de l'information (messagerie électronique et stockage de fichiers).

***
#### 5.2.1) To what extent are changes managed?

{.evaltext1}
* <span>Objectif</span> : The objective is to ensure that information security aspects are considered in case of any changes to the organization, business processes and IT systems (Change Management) in order to prevent these changes from causing an uncontrolled reduction in the information security level.
* <span>Référence</span> : IS0 27001: A.12.1.2
* <span>Maturité</span> : 1 ❌


**Exigences** : 

* *Sécurité des services de réseau* : Pour tous les services de réseau, les mécanismes de sécurité, les niveaux de service et les exigences de gestion, doivent être identifiés et intégrés dans les accords de services de réseau, que ces services soient fournis en interne ou externalisés (A.13.1.2).

**Observation**

Deux pare-feu Fortigate sont en cluster pour assurer la continuité de service en sortie du réseau interne (messagerie, SAP). Le réseau interne est en attente de rénovation avec l'arrivée de nouveau commutateur.

**Préconisation**

La cartographie du système d'information doit être réalisée afin de déterminer les parties critiques à redonder.

Les prérequis en termes de sécurité doivent avoir été déterminés et pris en compte dans composition du réseau afin de le protéger contre les menaces.

{.todo .msgblock}
Voir si on dispose d'un plan réseau

Des accords en termes de disponibilité doivent avoir été conclus avec les fournisseurs de services extérieurs, ou, ceux fournis dans le contrat avec le prestataire, doivent avoir été jugés acceptables au regard de la politique de sécurité.

#### 5.2.2) To what extent are development and testing environments separated from operational environments?

{.evaltext0}
* <span>Objectif</span> : The objective of separating the development, testing and operational environments is to ensure that the availability, confidentiality and integrity of productive data are maintained.
* <span>Référence</span> : IS0 27001: A.12.1.4
* <span>Maturité</span> : 0 ❌


**Exigences** :

* *Séparation des environnements de développement, de test et d’exploitation* : Il convient de séparer les environnements de développement, de test et d’exploitation pour réduire les risques d’accès ou de changements non autorisés dans l’environnement en exploitation (A.12.1.4).

**Observations**

Il est important de noter qu'il existe des environnements différents dans le domaine de l'informatique. Celui de production peut porter à confusion dans le domaine de l'industrie.

Environnement      | Utilisation
-------------------|-----------------------------------------------------------------------------------------------------------------
Test/Développement | Appareils (serveurs, réseau, clients) utilisés pour le développement d'applications ou des tests à portée restreinte (développeurs)
Validation         | Appareils servant à valider le bon fonctionnement d'une modification ou ajout avant de faire la bascule sur l'environnement de production
Production         | Appareils servant à faire fonctionner le système informatique utilisé au quotidien par les utilisateurs finaux.
Table: Les environnements informatiques

Il n'existe qu'un seul environnement présent sur le site.

**Préconisation**

Un environnement de validation devrait être présent pour valider que des changements n'impactent pas le fonctionnement des serveurs de production (cf [5.2.5](#toc_37)) ou le SMSI (accès non prévu à des données...)


#### 5.2.3) To what extent are IT systems protected against malware?

{.evaltext1}
* <span>Objectif</span> : The aim is to both technically and organizationally ensure the protection of IT systems against malware.
* <span>Référence</span> : IS0 27001: A.12.2
* <span>Maturité</span> : 1 ❌


**Exigence** :

* *Protection contre les logiciels malveillants* S’assurer que l’information et les moyens de traitement de l’information sont protégés contre les logiciels malveillants (A.12.2).
  * *Mesures contre les logiciels malveillants* : Il convient de mettre en œuvre des mesures de détection, de prévention et de
récupération, conjuguées à une sensibilisation des utilisateurs adaptée, pour se
protéger contre les logiciels malveillants (A.12.2.1).

**Observations** :

La protection contre les malwares repose sur l'utilisation seule d'un antivirus. Pas de mesures supplémentaires tel que la désactivation des protocoles obsolètes et activation des services réseaux inutiles.

Les serveurs regroupent un nombre de rôles ou d'applications important. Cela fragilise de fait le disponibilité de service en cas de problème logiciel ou matériel. Par exemple, sur l'unique contrôleur de domaine sont présents :

* un service de licence DASSAULT
* KÉLIO (+ un serveur Apache/Tomcat)
* un logiciel de gestion des incidents piloté par AI (OSRamp)
* des bases de données (Postgresql, Firebird)
* un logiciel de prise en main à distance (GOTO Assist) en attente de connexion
* un rôle de serveur de fichiers.

Le pare-feu n'est également pas activé sur le contrôleur de domaine. La configuration du serveur peut conduire un logiciel malveillant à prendre la main sur le système d'information, corrompre les sauvegardes et crée une interruption de service pour chacun des services présents sur le contrôleur.

**Préconisations** :

Le contrôleur de domaine étant un point sensible dans la sécurité du système d'information, il est donc recommandé de supprimer l'ensemble des applicatifs qui pourraient fonctionner sur un autre serveur. Cela permet de réduire la surface d'attaque et de garantir une meilleur continuité de service en éliminant ce point de défaillance individuel.

Une rupture technologique est également à prévoir pour les sauvegardes afin de préserver un jeu intact en cas attaque de rançongiciel qui toucherait VEEAM.

#### 5.2.4) To what extent are event logs recorded and analyzed?

{.evaltext0}
* <span>Objectif</span> : Event logs support the traceability of events in case of a security incident. This requires that events necessary to determine the causes are recorded and stored. In addition, the logging and analysis of activities in accordance with applicable legislation (e.g. Data Protection or Works Constitution Act) is required to determine which user account has made changes to IT systems.
* <span>Référence</span> : IS0 27001: A.12.4.1, A.12.4.2, A.12.4.3
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* *Journalisation des événements* :Il convient de créer, de tenir à jour et de revoir régulièrement les journaux d’événements enregistrant les activités de l’utilisateur, les exceptions, les défaillances et les événements liés à la sécurité de l’information.(A.12.4.1) 
* *Protection de l’information journalisée* : Il convient de protéger les moyens de journalisation et l’information journalisée
contre les risques de falsification ou d’accès non autorisé (A.12.4.2)
* *Journaux administrateur et opérateur* : Il convient de journaliser les activités de l’administrateur système et de l’opérateur système, ainsi que de protéger et de revoir régulièrement les journaux.(A.12.4.3)

**Observations**

La journalisation des événements est configurée par défaut. Pas de vérification sur ces journaux.

Les journaux ne sont pas protégés contre les risques de falsification.

Les activités des administrateurs ou comptes à pouvoir ne sont pas traités différemment des autres activités.

**Préconisations**

Des outils doivent être à disposition pour permettre l'exploitation des logs par des personnels autorisés

Les journaux systèmes doivent être transférés sur un système à part afin d'être protégés en cas d'intrusion.

Une solution doit être mise en place pour détecter les événements suspects.

#### 5.2.5) To what extent are vulnerabilities identified and addressed?

{.evaltext0}
* <span>Objectif</span> : Vulnerabilities increase the risk of IT systems being unable to meet the requirements for confidentiality, availability and integrity. Exploitation of vulnerabilities is among the possible ways for attackers to gain access to the IT system or to threaten its operating stability.
* <span>Référence</span> : IS0 27001: A.12.6
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* Gestion des vulnérabilités techniques (A.12.6.1)
* Restrictions liées à l’installation de logiciels (A.12.6.2)

**Observations** :

Pas de suivi des vulnérabilités des systèmes d'informations en exploitation. Seules les mises à jours Windows sont réalisées au fil de l'eau au gré des redémarrages serveurs. Il n'y a pas de politique de suivi des mises à jours, un correctif critique peut donc être appliqué plusieurs semaines après son déploiement.

**Préconisations**

Un suivi régulier des alertes du [CERT-FR](https://www.cert.ssi.gouv.fr/) (ou autre) doit être mis en place pour avoir une visibilité sur les alertes les plus urgentes. Des procédures doivent prévoir les modalités des mises à jour.

Dans l'idéal, un environnement de validation doit permettre de valider que la mise à jour n'a pas d'effets indésirables sur les serveurs de production. Par exemple, les correctifs de la faille [printNightmare a bénéficié de plusieurs correctifs](https://www.numerama.com/cyberguerre/739697-microsoft-met-enfin-un-terme-a-printnightmare-son-cauchemar-de-lete.html) ayant eu plusieurs fois un impact sur le fonctionnement des imprimantes.

Une liste des logiciels autorisés, à l'instar de [1.3.3](#toc_10), doit être définie.

Certains logiciels requiert une élévation de privilèges et nécessite que de droits d'administration soient accordés à un utilisateur, il faudra trouver dans ce cas une solution plus restrictive.
 

#### 5.2.6) To what extent are IT systems technically checked (system audit)?

{.evaltext0}
* <span>Objectif</span> : The objective of technical checks is the detection of states which can jeopardize the availability, confidentiality or integrity of IT systems.
* <span>Référence</span> : IS0 27001: A12.7, A.18.2.3
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* Les exigences et activités d’audit impliquant des vérifications sur des systèmes en exploitation doivent être prévues avec soin et validées afin de réduire au minimum les perturbations subies par les processus métier (A.12.7).
* Vérification de la conformité technique (A.18.2.3)

Des audits doivent être conduits régulièrement afin de déterminer si le système d'information est conforme à la politique de sécurité ou si les mesures d'améliorations ou de correction sont effectives.

Les auditeurs doivent être formés à réaliser ces audits.

#### 5.2.7) To what extent is the network of the organization managed?

{.evaltext0}
* <span>Objectif</span> : IT systems in a network are exposed to different risks or have different protection needs. In order to detect or prevent unintended data exchange or access between these IT systems, they are subdivided into suitable segments and access is controlled and monitored by means of security technologies.
* <span>Référence</span> : IS0 27001: A.13.1.1, A.13.1.3
* <span>Maturité</span> : 0 ❌


**Exigences** : 

* Les réseaux doivent être gérés et contrôlés pour protéger l’information contenue dans les systèmes et les applications (A.13.1.1). 
* Pour tous les services de réseau, les mécanismes de sécurité, les niveaux de service et les exigences de gestion, doivent être identifiés et intégrés dans les accords de services de réseau, que ces services soient fournis en interne ou externalisés (A.13.1.2)

**Observation**

Pas de ségrégation en place sur le réseau. La différenciation entre les postes clients et les serveurs est uniquement réalisée sur l'adresse du réseau : 

* 10.0.80.x pour les serveurs et appareils réseaux.
* 10.0.81.x pour les postes clients

La mise en place de la nouvelle infrastructure (Fortigate + commutateurs réseaux récents) va permettre de segmenter le réseau informatique en VLANs et restreindre les accès.

Surveillance du réseau :

Le pare-feu Fortigate permet de monitorer le réseau sur un temps limité (en heures). 

**Préconisations**

Mettre en place des cloisenements : Les systèmes extérieurs non contrôlés (par exemple, fifteen, badgeuses) doivent être séparés des réseau d'exploitation d'AGLA FORM.

Une solution de type [FortiAnalyzer](https://www.fortinet.com/fr/products/management/fortianalyzer) doit être prévue afin de pouvoir conserver les traces d'activité réseau sur une plus longue durée (en jours, voire mois, selon licence et stockage).

***
#### 5.3.1) To what extent is information security considered in new or further developed IT systems?

{.evaltext0}
* <span>Objectif</span> : Information security is an integral part of the entire lifecycle of IT systems. This particularly includes consideration of information security requirements in the development or acquisition of IT systems.
* <span>Référence</span> : IS0 27001: A.14.1
* <span>Maturité</span> : 0 ❌


**Exigences** :

* *Exigences de sécurité applicables aux systèmes d’information* : Veiller à ce que la sécurité de l’information fasse partie intégrante des systèmes d’information tout au long de leur cycle de vie. Cela inclut notamment des exigences spécifiques pour les systèmes d’information fournissant des services sur les réseaux publics. (A.14.1)
  * *Analyse et spécification des Exigences de sécurité de l’information* : Il convient que les exigences liées à la sécurité de l’information figurent dans les exigences des nouveaux systèmes d’information ou des changements apportés aux systèmes existants (A.14.1.1).
  * *Sécurisation des services d’application sur les réseaux publics* : Il convient de protéger l’information liée aux services d’application transmise sur les réseaux publics contre les activités frauduleuses, les différends contractuels, ainsi que la divulgation et la modification non autorisées (A.14.1.2).
  * *Protection des transactions liées aux services d’application* : Il convient de protéger l’information impliquée dans les transactions liées aux services d’application pour empêcher une transmission incomplète, des erreurs d’acheminement, la modification non autorisée, la divulgation non autorisée, la duplication non autorisée du message ou sa réémission (A.14.1.3)

**Observation**

Idem [1.2.3](#toc_6). 

L'absence d'une politique de sécurité au niveau de la société ne permet pas la considération d'une politique de sécurité.

#### 5.3.2) To what extent are requirements for network services defined?

{.evaltext0}
* <span>Objectif</span> : Network services have different requirements for information security, quality of data transfer or management. It is important to know these criteria and the scope of use of the different network services.
* <span>Référence</span> : IS0 27001: A.13.1.2
* <span>Maturité</span> : 0 ❌

**Exigences**

* *Sécurité des services de réseau* : Pour tous les services de réseau, il convient d’identifier les mécanismes de sécurité, les niveaux de service et les exigences de gestion, et de les intégrer dans les accords de services de réseau, que ces services soient fournis en interne ou externalisés.

#### 5.3.3) To what extent is the return and secure removal of information assets from external IT services regulated?

{.evaltext0}
* <span>Objectif</span> : In order to ensure control over the information assets as the information owner, it is necessary that the information assets can be safely removed or are returned, if required, when terminating the IT service.
* <span>Référence</span> : IS0 27017: CLD.8.1.5
* <span>Maturité</span> : 0 ❌

Relatif à la gestion des périphériques [3.1.4) To what extent is the handling of mobile IT devices and mobile data storage devices managed?](#toc_24)

**Exigences**

* Les actifs du client du service cloud qui se trouvent dans les locaux du fournisseur de services cloud doivent être supprimés et retournés si nécessaire, en temps opportun à la résiliation du contrat de service cloud (ISO 27017 CLD.8.1.5).


{.todo .msgblock}
Voir avec Marc comment sont gérés les actifs sur le cloud

#### 5.3.4) To what extent is information protected in shared external IT services?

{.evaltext0}
* <span>Objectif</span> : Clear segregation between individual clients must be ensured such as to protect own information in external IT services at all times and to prevent it from being accessed by other organizations (clients).
* <span>Référence</span> : ISO 27017: CLD.9.5.1, CLD.9.5.2
* <span>Maturité</span> : 0 ❌

**Exigences**

* L’environnement virtuel d’un client de service cloud s’exécutant sur un service cloud doit être protégé contre les autres clients de service cloud et les personnes non autorisées (ISO 27017 CLD.9.5.1).
* Les machines virtuelles dans un environnement de cloud computing doivent être renforcées pour répondre aux besoins de l’entreprise ((ISO 27017 CLD.9.5.1)

**Observation**

Pas d'utilisation de service cloud en tant que PAAS.

***
### 6) Supplier Relationships
#### 6.1.1) To what extent is information security ensured among contractors and cooperation partners?

{.evaltext0}
* <span>Objectif</span> : An appropriate level of information security is also maintained while collaborating with cooperation partners and contractors.
* <span>Référence</span> : ISO 27001: A.15.1, A.15.2.1
* <span>Maturité</span> : 0 ❌


*  *Sécurité de l’information dans les relations avec les fournisseurs*  : Garantir la protection des actifs de l’organisation accessibles aux fournisseurs (A.15.1).
  * *Politique de sécurité de l'information dans les relations avec les fournisseurs* : Il convient de convenir avec le fournisseur les exigences de sécurité de l’information pour limiter les risques résultant de l’accès du fournisseur aux actifs de l’organisation et de les documenter (A.15.1.1).
  * *La sécurité dans les accords conclus avec les fournisseurs* : Il convient que les exigences applicables liées à la sécurité de l’information soient établies et convenues avec chaque fournisseur pouvant avoir accès, traiter, stocker, communiquer ou fournir des composants de l’infrastructure informatique destinés à l’information de l’organisation (A.15.1.2).
  * *Chaine d’approvisionnement informatique* : Il convient que les accords conclus avec les fournisseurs incluent des exigences
sur le traitement des risques de sécurité de l’information associés à la chaîne d’approvisionnement des produits et des services informatiques.

{.todo .msgblock}
Voir avec Serge pour les relations avec les fournisseurs

**Préconisations**

En cas de sous-traitance à des tiers, s’assurer que : 

* les obligations en matière de traitement de données à caractère personnel sont fixées dans un contrat.
* les conditions relatives à la sécurité de l'information et à la protection des données à caractère personnel font l’objet d’un accord avec les tiers et sont documentées afin de réduire les risques relatifs à l’accès des tiers aux moyens d’information
* qu'une politique de « sécurité de la sous-traitance à des tiers » est en place et s'assurer que les directives décrites sont respectées par les tiers.
* Un contrat de sous-traitance est conclu avec le fournisseur selon les dispositions du RGPD.
* Il est fixé dans un contrat ce qu’il faut faire avec les données lorsque la collaboration prend fin.

Document : Code de conduite des partenaires Valeo  22 10 2018 FR.pdf

#### 6.1.2) To what extent is non-disclosure regarding the exchange of information contractually agreed?

{.evaltext0}
* <span>Objectif</span> : Non-disclosure agreements provide legal protection of an organization’s information particularly where information is exchanged beyond the boundaries of the organization.
* <span>Référence</span> : IS0 27001: A.13.2.2, A.13.2.4
* <span>Maturité</span> : 0 ❌

**Exigences** :

* *Accords en matière de transfert d’information* : Il convient que les accords traitent du transfert sécurisé de l’information liée à
l’activité entre l’organisation et les tiers. (A.13.2.2)
* *Engagements de confidentialité ou de non-divulgation* : Il convient d’identifier, de revoir régulièrement et de documenter les exigences en matière d’engagements de confidentialité ou de non-divulgation, conformément aux besoins de l’organisation en matière de protection de l’information (A.13.2.4).

Les exigences en matière d’engagements de confidentialité ou de non-divulgation, doivent être identifiées, vérifiées régulièrement et documentées conformément aux besoins de l’organisme.

Des accords doivent traiter du transfert sécurisé de l’information liée à l’activité entre l’organisme et les tiers (ISO A.13.2.2) :

* Accords traitant du transfert sécurisé de l’information liée à l’activité
* Identification des responsabilités de gestion, de la répartition et de la réception de l’information
* Procédure de gestion de la traçabilité et la non-répudiation
* Identification des obligations et des responsabilités des uns et des autres en cas d’incident lié à la sécurité de l’information
* Rapports de traitement des incidents liés à la sécurité de l’information

***
### 7) Compliance
#### 7.1.1) To what extent is compliance with regulatory and contractual provisions ensured?

{.evaltext0}
* <span>Objectif</span> : Non-compliance with legal, regulatory or contractual provisions can create risks to the information security of customers and the own organization. Therefore, it is essential to ensure that these provisions are known and observed.
* <span>Référence</span> : IS0 27001: A.18.1.1, A.18.1.2, A.18.1.3, A.18.1.5
* <span>Maturité</span> : 0 ❌

**Exigences**

* *Identification de la législation et des exigences contractuelles applicables* : Il convient, pour chaque système d’information et pour l’organisation elle-même, de définir, documenter et mettre à jour explicitement toutes les exigences légales, réglementaires et contractuelles en vigueur, ainsi que l’approche adoptée par l’organisation pour satisfaire à ces exigences (A.18.1.1).
* *Droits de propriété intellectuelle* : Il convient de mettre en œuvre des procédures appropriées visant à garantir la conformité avec les exigences légales, réglementaires et contractuelles relatives aux droits de propriété intellectuelle et à l’utilisation de logiciels propriétaires (A.18.1.2).
* *Protection des enregistrements* : Il convient de protéger les enregistrements de la perte, de la destruction, de la falsification, des accès non autorisés et des diffusions non autorisées conformément aux exigences légales, réglementaires, contractuelles et aux
exigences métier (A.18.1.3).
* *Réglementation relative aux mesures cryptographiques* : Il convient de prendre des mesures cryptographiques conformément aux accords, lois et réglementations applicables (A.18.1.5).

S'assurer que :

* la situation est conforme à la sécurité de l'information et à la protection des données à caractère personnel telle que décrite dans les politiques.
* il n'y a pas violation de la législation et des obligations contractuelles, statutaires, réglementaires ou légales relatives à la sécurité de l'information et à la protection des données à caractère personnel.
* les moyens de garantir la sécurité de l'information et la protection des données ont été mises en œuvre et
sont opérationnelles conformément aux attentes de la direction.
* un processus disciplinaire formel est en place pour les travailleurs ayant commis une infraction à la sécurité de l'information ou à la protection des données à caractère personnel.

Il faut s'assurer que les moyens cryptographiques en œuvre sont conformes à la législation en vigueur dans le pays où elle est utilisée.

#### 7.1.2) To what extent is the protection of personally identifiable data taken into account when implementing information security?

{.evaltext0}
* <span>Objectif</span> : Privacy and protection of personally identifiable data are taken into account in the implementation of information security as required by relevant national legislation and regulations, where applicable.
* <span>Référence</span> : IS0 27001: A.18.1.4
* <span>Maturité</span> : 0 ❌

**Exigence**

* *Protection de la vie privée et protection des données à caractère personnel* : Il convient de garantir la protection de la vie privée et la protection des données à caractère personnel telles que l’exigent la législation et les réglementations applicables, le cas échéant.

Voir chapitre RGPD.

***
## Prototype Protection
### 8) Prototype Protection
#### 8.1.1) To what extent is a security concept available describing minimum requirements regarding the physical and environmental security for prototype protection?

{.evaltext0}
* <span>Objectif</span> : The necessary measures for prototype protection must be applied to and implemented on properties and facilities of suppliers, development partners and service providers. A security concept must be established by the respective operator. Implementation and observation of the physical and environmental security measures defined in the security concept must be ensured by the responsible operator.


**Observation**

Pas de mesure de sécurité prises pour protéger des prototypes présents sur site.

**Préconisations**

* une restriction des accès physique aux lieux où sont présent les prototypes.
* une politique concernant la diffusion d'information (interdiction appareils photo, clause de non divulgation...)

{.todo .msgblock}
Voir avec Serge ce qui est présent sur place

#### 8.1.2) To what extent is perimeter security existent preventing unauthorized access to protected property objects?

{.evaltext0}
* <span>Objectif</span> : Unauthorized access to properties where vehicles, components or parts classified as requiring protection are manufactured, processed or stored shall be prevented.

Non évalué
#### 8.1.3) To what extent is the outer skin of the protected buildings constructed such as to prevent removal or opening of outer-skin components using standard tools?

{.evaltext0}
* <span>Objectif</span> : Unauthorized access to buildings/security areas where vehicles, components or parts classified as requiring protection are manufactured, processed or stored shall be prevented.

Non évalué
#### 8.1.4) It must be ensured, that the camouflage regulations are known to each project member and observed in order to guarantee adequate view protection of trial vehicles.

{.evaltext0}
* <span>Objectif</span> : It must be ensured that unauthorized viewing of vehicles, components or parts classified as requiring protection is prevented.

Non évalué
#### 8.1.5) To what extent is the protection against unauthorized entry regulated in the form of access control?

{.evaltext0}
* <span>Objectif</span> : It must be ensured that all points of access to security areas where vehicles, components or parts classified as requiring protection are manufactured, processed or stored are protected against unauthorized entry by adequate measures.

Non évalué
#### 8.1.6) To what extent are the premises to be secured monitored for intrusion?

{.evaltext0}
* <span>Objectif</span> : It must be ensured that premises where vehicles, components or parts classified as requiring protection are manufactured, processed or stored are monitored for intrusion. Timely alarm processing is ensured.

Non évalué
#### 8.1.7) To what extent is a documented visitor management in place?

{.evaltext0}
* <span>Objectif</span> : Protection against unauthorized access to security areas where vehicles, components or parts classified as requiring protection are manufactured, processed or stored, including traceable documentation.

Non évalué
#### 8.1.9) To what extent is on-site client segregation existent?

{.evaltext0}
* <span>Objectif</span> : In order to ensure protection of the client-specific know-how at all times, a clear segregation of clients must be ensured. This particularly involves protection against unauthorized viewing and access to areas where vehicles, components or parts classified as requiring protection are processed or stored.

Non évalué
***
#### 8.2.1) To what extent are non-disclosure agreements/obligations existent according to the valid contractual law?

{.evaltext0}
* <span>Objectif</span> : When transmitting information classified as requiring protection, it must be ensured that external organizations are obliged to meet the information security requirements and that the related necessary measures are implemented. The necessary legal basis for this obligation is provided by non-disclosure agreements. Hence, it must be ensured that information classified as requiring protection is transmitted only if such a non-disclosure agreement has been entered and is legally effective.

Non évalué
#### 8.2.2) To what extent are requirements for commissioning subcontractors known and fulfilled?

{.evaltext0}
* <span>Objectif</span> : When involving subcontractors, the minimum requirements for prototype protection must be met.

Non évalué
#### 8.2.3) To what extent do employees and project members evidently participate in training and awareness measures regarding the handling of prototypes?

{.evaltext0}
* <span>Objectif</span> : In trainings/awareness seminars on the subject of prototype protection, employees must obtain the necessary knowledge and skills for a security-conscious handling of vehicles, components and parts classified as requiring protection.

Non évalué
#### 8.2.4) To what extent are security classifications of the project and the resulting security measures known?

{.evaltext0}
* <span>Objectif</span> : It must be ensured that the security classification and requirements in relation to the project progress are known to and observed by each project member.

Non évalué
#### 8.2.5) To what extent is a process defined for granting access to security areas?

{.evaltext0}
* <span>Objectif</span> : A process is defined for the protection against unauthorized access to security areas where vehicles, components or parts classified as requiring protection are manufactured, processed or stored.

Non évalué
#### 8.2.6) To what extent are regulations for image recording and handling of created image material existent?

{.evaltext0}
* <span>Objectif</span> : Regulations for recording images of vehicles, components or parts classified as requiring protection must be defined in order to prevent unauthorized creation or transmission of such image material.

Non évalué
#### 8.2.7) To what extent is a process for carrying along and using mobile video and photography devices in(to) defined security areas established?

{.evaltext0}
* <span>Objectif</span> : A process is defined for carrying along and using mobile video and photography devices in(to) security areas where vehicles, components or parts classified as requiring protection are manufactured, processed or stored. Unauthorized creation or transmission of image material must be prevented.

Non évalué
***
#### 8.3.1) To what extent are transports of vehicles, components or parts classified as requiring protection arranged according to the customer requirements?

{.evaltext0}
* <span>Objectif</span> : While being transported, vehicles, components and parts classified as requiring protection must be protected against unauthorized viewing, unauthorized image recording and access.

Non évalué
#### 8.3.2) To what extent is it ensured that vehicles, components and parts classified as requiring protection are parked/stored in accordance with customer requirements?

{.evaltext0}
* <span>Objectif</span> : While being parked/stored, vehicles, components and parts classified as requiring protection must be protected against unauthorized viewing, unauthorized photography and access.

Non évalué
***
#### 8.4.1) To what extent are the predefined camouflage regulations implemented by the project members?

{.evaltext0}
* <span>Objectif</span> : It must be ensured, that the camouflage regulations are known to each project member and observed in order to guarantee adequate view protection of trial vehicles.

Non évalué
#### 8.4.2) To what extent are measures for protecting approved test and trial grounds observed/implemented?

{.evaltext0}
* <span>Objectif</span> : In order to maintain an undisturbed and secured trial operation on test and trial grounds, the respective protective measures defined by the customer must be observed.

Non évalué
#### 8.4.3) To what extent are protective measures for approved test and trial drives in public observed/implemented?

{.evaltext0}
* <span>Objectif</span> : It must be ensured that the respective customer requirements for the operation of trial vehicles classified as requiring protection on public roads are known and observed.

Non évalué
***
#### 8.5.1) To what extent are security requirements for presentations and events involving vehicles, components or parts classified as requiring protection known?

{.evaltext0}
* <span>Objectif</span> : It must be ensured that the respective customer-specific security requirements for presentations and events involving vehicles, components or parts classified as requiring protection are known.

Non évalué
#### 8.5.2) To what extent are the protective measures for film and photo shootings involving vehicles, components or parts classified as requiring protection known?

{.evaltext0}
* <span>Objectif</span> : It must be ensured that the respective customer-specific security requirements for film and photo shootings involving vehicles, components or parts classified as requiring protection are known.

Non évalué
***
## Data Protection
### 9) Data Protection
#### 9.1) To what extent is the implementation of data protection organized?

{.evaltext0}


**Observations**

Le RGPD impose le principe de "protection des données dès la conception et par défaut". 

Pas de registre des activités de traitements, ni d'inventaire sur le traitement des données personnelles (fournisseurs, clients, salariés).

Il n'y a pas de conformité avec le RGPD.

{.warning .center .msgblock}
Le règlement général de la protection des données personnelles (RGPD) prévoit **une sanction pouvant aller jusqu'à 4% du CA annuel** en cas de manquement au droit des personnes.
 
**Préconisation**

Pour la mise en conformité avec le RGPD:

* Faire l'inventaire des fichiers de données contenant des données personnelles.
* [Mettre en place un registre de traitements](https://www.cnil.fr/fr/RGDP-le-registre-des-activites-de-traitement). 
* Limiter l'accès aux données (mail, téléphone...) dans le **seul but déclaré** du traitement et aux **seules personnes déclarées** pour le traitement.

Un [responsable de traitement](https://www.cnil.fr/fr/definition/responsable-de-traitement) doit-être désigné. Il est en charge du bon respect du (RGPD)[https://www.cnil.fr/fr/reglement-europeen-protection-donnees/].

Aucune donnée personnelle ne doit être présente dans le dossier commun (partage `public`) de la société.

Ressources :

* [Check-list RGPD pour TPE/PME](https://www.cnil.fr/sites/default/files/atoms/files/check-list_rgpd_pour_les_tpe-pme.pdf).
* [Guide pratique de sensibilisation au RGPD](https://www.cnil.fr/sites/default/files/atoms/files/bpi-cnil-rgpd_guide-tpe-pme.pdf)
* [Guide de la sécurité des données personnelles](https://www.cnil.fr/sites/default/files/atoms/files/cnil_guide_securite_personnelle.pdf)


#### 9.2) To what extent are organizational measures taken in order to ensure that personally identifiable data is processed in conformance with legislation?

{.evaltext0}

Non évalué
#### 9.3) To what extent is it ensured that the internal processes or workflows are carried out according to the currently valid data protection regulations and that these are regularly subjected to a quality check?

{.evaltext0}

Non évalué
#### 9.4) To what extent are the relevant processing procedures documented with regard to their admissibility according to data protection law?

{.evaltext0}

Non évalué
***
## Synthèse

### Information Security

 ISA | Question | Maturité
-----|----------|---------
1.1.1 | To what extent are information security policies available? | 0 ❌
1.2.1 | To what extent is information security managed within the organization? | 0 ❌
1.2.2 | To what extent are information security responsibilities organized? | 0 ❌
1.2.3 | To what extent are information security requirements taken into account in projects? | 0 ❌
1.2.4 | To what extent are the responsibilities between external IT service providers and the own organization defined? | -
1.3.1 | To what extent are information assets identified and recorded? | 1 ❌
1.3.2 | To what extent are information assets classified and managed in terms of their protection needs? | 0 ❌
1.3.3 | To what extent is it ensured that only evaluated and approved external IT services are used for processing the organization’s information assets? | 0 ❌
1.4.1 | To what extent are information security risks managed? | 0 ❌
1.5.1 | To what extent is compliance with information security ensured in procedures and processes? | 0 ❌
1.5.2 | To what extent is the ISMS reviewed by an independent authority? | 0 ❌
1.6.1 | To what extent are information security events processed? | 0 ❌
2.1.1 | To what extent is the qualification of employees for sensitive work fields ensured? | -
2.1.2 | To what extent is all staff contractually bound to comply with information security policies? | 1 ❌
2.1.3 | To what extent is staff made aware of and trained with respect to the risks arising from the handling of information? | 1 ❌
2.1.4 | To what extent is teleworking regulated? | 0 ❌
3.1.1 | To what extent are security zones managed to protect information assets? | 0 ❌
3.1.2 | To what extent is information security ensured in exceptional situations? | 1 ❌
3.1.3 | To what extent is the handling of supporting assets managed? | 0 ❌
3.1.4 | To what extent is the handling of mobile IT devices and mobile data storage devices managed? | 0 ❌
4.1.1 | To what extent is the use of identification means managed? | 1 ❌
4.1.2 | To what extent is the user access to network services, IT systems and IT applications secured? | 1 ❌
4.1.3 | To what extent are user accounts and login information securely managed and applied? | 1 ❌
4.2.1 | To what extent are access rights assigned and managed? | -
5.1.1 | To what extent is the use of cryptographic procedures managed? | 0 ❌
5.1.2 | To what extent is information protected during transfer? | 0 ❌
5.2.1 | To what extent are changes managed? | 1 ❌
5.2.2 | To what extent are development and testing environments separated from operational environments? | 0 ❌
5.2.3 | To what extent are IT systems protected against malware? | 1 ❌
5.2.4 | To what extent are event logs recorded and analyzed? | 0 ❌
5.2.5 | To what extent are vulnerabilities identified and addressed? | 0 ❌
5.2.6 | To what extent are IT systems technically checked (system audit)? | 0 ❌
5.2.7 | To what extent is the network of the organization managed? | 0 ❌
5.3.1 | To what extent is information security considered in new or further developed IT systems? | 0 ❌
5.3.2 | To what extent are requirements for network services defined? | -
5.3.3 | To what extent is the return and secure removal of information assets from external IT services regulated? | -
5.3.4 | To what extent is information protected in shared external IT services? | -
6.1.1 | To what extent is information security ensured among contractors and cooperation partners? | -
6.1.2 | To what extent is non-disclosure regarding the exchange of information contractually agreed? | -
7.1.1 | To what extent is compliance with regulatory and contractual provisions ensured? | -
7.1.2 | To what extent is the protection of personally identifiable data taken into account when implementing information security? | -

### Prototype Protection

 ISA | Question | Maturité
-----|----------|---------
8.1.1 | To what extent is a security concept available describing minimum requirements regarding the physical and environmental security for prototype protection? | 0 ❌
8.1.2 | To what extent is perimeter security existent preventing unauthorized access to protected property objects? | -
8.1.3 | To what extent is the outer skin of the protected buildings constructed such as to prevent removal or opening of outer-skin components using standard tools? | -
8.1.4 | It must be ensured, that the camouflage regulations are known to each project member and observed in order to guarantee adequate view protection of trial vehicles. | -
8.1.5 | To what extent is the protection against unauthorized entry regulated in the form of access control? | -
8.1.6 | To what extent are the premises to be secured monitored for intrusion? | -
8.1.7 | To what extent is a documented visitor management in place? | -
8.1.9 | To what extent is on-site client segregation existent? | -
8.2.1 | To what extent are non-disclosure agreements/obligations existent according to the valid contractual law? | -
8.2.2 | To what extent are requirements for commissioning subcontractors known and fulfilled? | -
8.2.3 | To what extent do employees and project members evidently participate in training and awareness measures regarding the handling of prototypes? | -
8.2.4 | To what extent are security classifications of the project and the resulting security measures known? | -
8.2.5 | To what extent is a process defined for granting access to security areas? | -
8.2.6 | To what extent are regulations for image recording and handling of created image material existent? | -
8.2.7 | To what extent is a process for carrying along and using mobile video and photography devices in(to) defined security areas established? | -
8.3.1 | To what extent are transports of vehicles, components or parts classified as requiring protection arranged according to the customer requirements? | -
8.3.2 | To what extent is it ensured that vehicles, components and parts classified as requiring protection are parked/stored in accordance with customer requirements? | -
8.4.1 | To what extent are the predefined camouflage regulations implemented by the project members? | -
8.4.2 | To what extent are measures for protecting approved test and trial grounds observed/implemented? | -
8.4.3 | To what extent are protective measures for approved test and trial drives in public observed/implemented? | -
8.5.1 | To what extent are security requirements for presentations and events involving vehicles, components or parts classified as requiring protection known? | -
8.5.2 | To what extent are the protective measures for film and photo shootings involving vehicles, components or parts classified as requiring protection known? | -

### Data Protection

 ISA | Question | Maturité
-----|----------|---------
9.1 | To what extent is the implementation of data protection organized? | 0 ❌
9.2 | To what extent are organizational measures taken in order to ensure that personally identifiable data is processed in conformance with legislation? | -
9.3 | To what extent is it ensured that the internal processes or workflows are carried out according to the currently valid data protection regulations and that these are regularly subjected to a quality check? | -
9.4 | To what extent are the relevant processing procedures documented with regard to their admissibility according to data protection law? | -

