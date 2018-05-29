-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le :  mar. 29 mai 2018 à 12:06
-- Version du serveur :  5.7.19
-- Version de PHP :  7.1.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données :  `alea_data_est`
--
DROP DATABASE IF EXISTS `alea_data_est`;
CREATE DATABASE IF NOT EXISTS `alea_data_est` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin;
USE `alea_data_est`;

-- --------------------------------------------------------

--
-- Structure de la table `champ`
--

CREATE TABLE `champ` (
  `id` int(10) UNSIGNED NOT NULL,
  `ressource_id` int(10) UNSIGNED DEFAULT NULL,
  `clef` varchar(255) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Déchargement des données de la table `champ`
--

INSERT INTO `champ` (`id`, `ressource_id`, `clef`) VALUES
(1, 1, 'nom'),
(2, 1, 'prenom'),
(3, 1, 'age'),
(4, 1, 'sexe'),
(5, 1, 'ville'),
(6, NULL, 'truc'),
(7, NULL, 'truc'),
(9, NULL, 'machin'),
(10, 6, 'rule1'),
(11, 6, 'rule2'),
(12, 6, 'rule3'),
(13, 6, 'rule4'),
(14, 6, 'rule5'),
(15, 6, 'rule6'),
(16, 6, 'rule7'),
(17, 6, 'rule8'),
(18, 6, 'rule9'),
(19, 6, 'rule10'),
(20, 6, 'rule11');

-- --------------------------------------------------------

--
-- Structure de la table `champ_parametre`
--

CREATE TABLE `champ_parametre` (
  `id` int(10) UNSIGNED NOT NULL,
  `champ_id` int(10) UNSIGNED NOT NULL,
  `regle_parametre_id` int(10) UNSIGNED NOT NULL,
  `valeur` varchar(255) COLLATE utf8_bin DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Déchargement des données de la table `champ_parametre`
--

INSERT INTO `champ_parametre` (`id`, `champ_id`, `regle_parametre_id`, `valeur`) VALUES
(8, 1, 8, 'nom'),
(9, 2, 8, 'prenom'),
(10, 3, 10, '100'),
(11, 4, 1, '[Homme|Femme]'),
(12, 5, 8, 'ville'),
(13, 3, 9, '0'),
(18, 9, 9, NULL),
(19, 9, 10, NULL),
(21, 10, 1, 'abc|def'),
(22, 11, 2, '12'),
(23, 12, 3, '25'),
(24, 13, 4, '25'),
(25, 14, 5, '45'),
(26, 15, 6, '10'),
(27, 16, 12, '10'),
(28, 17, 11, '10'),
(29, 18, 7, '10'),
(30, 19, 8, 'Prenom'),
(31, 20, 9, '10'),
(32, 20, 10, '24');

-- --------------------------------------------------------

--
-- Structure de la table `parametre`
--

CREATE TABLE `parametre` (
  `id` int(10) UNSIGNED NOT NULL,
  `nom` varchar(255) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Déchargement des données de la table `parametre`
--

INSERT INTO `parametre` (`id`, `nom`) VALUES
(1, 'boolean'),
(2, 'string'),
(3, 'int'),
(4, 'float');

-- --------------------------------------------------------

--
-- Structure de la table `regle`
--

CREATE TABLE `regle` (
  `id` int(10) UNSIGNED NOT NULL,
  `nom` varchar(255) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Déchargement des données de la table `regle`
--

INSERT INTO `regle` (`id`, `nom`) VALUES
(1, 'Regex'),
(2, 'Inférieur'),
(3, 'Supérieur'),
(4, 'Inférieur ou égal'),
(5, 'Supérieur ou égal'),
(6, 'Egal'),
(7, 'Pair'),
(8, 'Impair'),
(9, 'Multiple de'),
(10, 'Dictionnaire'),
(11, 'Compris entre');

-- --------------------------------------------------------

--
-- Structure de la table `regle_parametre`
--

CREATE TABLE `regle_parametre` (
  `id` int(10) UNSIGNED NOT NULL,
  `regle_id` int(10) UNSIGNED NOT NULL,
  `parametre_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Déchargement des données de la table `regle_parametre`
--

INSERT INTO `regle_parametre` (`id`, `regle_id`, `parametre_id`) VALUES
(1, 1, 2),
(2, 2, 3),
(3, 3, 3),
(4, 4, 3),
(5, 5, 3),
(6, 6, 3),
(7, 9, 4),
(8, 10, 2),
(9, 11, 4),
(10, 11, 4),
(11, 8, NULL),
(12, 7, NULL);

-- --------------------------------------------------------

--
-- Structure de la table `ressource`
--

CREATE TABLE `ressource` (
  `id` int(10) UNSIGNED NOT NULL,
  `nom` varchar(255) COLLATE utf8_bin NOT NULL,
  `createur` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `date_creation` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Déchargement des données de la table `ressource`
--

INSERT INTO `ressource` (`id`, `nom`, `createur`, `date_creation`) VALUES
(1, 'user', '2alheure', '2018-05-20 13:27:38'),
(2, 'machin', 'frambur', '2018-05-20 15:34:43'),
(3, 'truc', '2dtension', '2018-05-20 15:34:43'),
(6, 'test', 'test', '2018-05-29 11:33:57'),
(7, 'test_generate', '2alheure', '2018-05-29 12:05:50');

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `champ`
--
ALTER TABLE `champ`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ressource_id` (`ressource_id`);

--
-- Index pour la table `champ_parametre`
--
ALTER TABLE `champ_parametre`
  ADD PRIMARY KEY (`id`),
  ADD KEY `champ_id` (`champ_id`),
  ADD KEY `regle_parametre_id` (`regle_parametre_id`);

--
-- Index pour la table `parametre`
--
ALTER TABLE `parametre`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `regle`
--
ALTER TABLE `regle`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `regle_parametre`
--
ALTER TABLE `regle_parametre`
  ADD PRIMARY KEY (`id`),
  ADD KEY `regle_id` (`regle_id`),
  ADD KEY `parametre_id` (`parametre_id`);

--
-- Index pour la table `ressource`
--
ALTER TABLE `ressource`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nom` (`nom`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `champ`
--
ALTER TABLE `champ`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT pour la table `champ_parametre`
--
ALTER TABLE `champ_parametre`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT pour la table `parametre`
--
ALTER TABLE `parametre`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT pour la table `regle`
--
ALTER TABLE `regle`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT pour la table `regle_parametre`
--
ALTER TABLE `regle_parametre`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT pour la table `ressource`
--
ALTER TABLE `ressource`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `champ`
--
ALTER TABLE `champ`
  ADD CONSTRAINT `champ_ibfk_1` FOREIGN KEY (`ressource_id`) REFERENCES `ressource` (`id`);

--
-- Contraintes pour la table `champ_parametre`
--
ALTER TABLE `champ_parametre`
  ADD CONSTRAINT `champ_parametre_ibfk_1` FOREIGN KEY (`champ_id`) REFERENCES `champ` (`id`),
  ADD CONSTRAINT `champ_parametre_ibfk_2` FOREIGN KEY (`regle_parametre_id`) REFERENCES `regle_parametre` (`id`);

--
-- Contraintes pour la table `regle_parametre`
--
ALTER TABLE `regle_parametre`
  ADD CONSTRAINT `regle_parametre_ibfk_1` FOREIGN KEY (`regle_id`) REFERENCES `regle` (`id`),
  ADD CONSTRAINT `regle_parametre_ibfk_2` FOREIGN KEY (`parametre_id`) REFERENCES `parametre` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
