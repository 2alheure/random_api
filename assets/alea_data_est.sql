-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le :  ven. 11 mai 2018 à 13:39
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

DROP TABLE IF EXISTS `champ`;
CREATE TABLE IF NOT EXISTS `champ` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ressource_id` int(10) UNSIGNED NOT NULL,
  `type_id` int(10) UNSIGNED NOT NULL,
  `clef` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ressource_id` (`ressource_id`),
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Structure de la table `parametre`
--

DROP TABLE IF EXISTS `parametre`;
CREATE TABLE IF NOT EXISTS `parametre` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `nom` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

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

DROP TABLE IF EXISTS `regle`;
CREATE TABLE IF NOT EXISTS `regle` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `nom` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

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
(9, 'Multiple de');

-- --------------------------------------------------------

--
-- Structure de la table `regle_parametre`
--

DROP TABLE IF EXISTS `regle_parametre`;
CREATE TABLE IF NOT EXISTS `regle_parametre` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_regle` int(10) UNSIGNED NOT NULL,
  `id_parametre` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_regle` (`id_regle`),
  KEY `id_parametre` (`id_parametre`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Déchargement des données de la table `regle_parametre`
--

INSERT INTO `regle_parametre` (`id`, `id_regle`, `id_parametre`) VALUES
(1, 1, 1),
(2, 2, 3),
(3, 3, 3),
(4, 4, 3),
(5, 5, 3),
(6, 6, 3),
(7, 9, 2);

-- --------------------------------------------------------

--
-- Structure de la table `ressource`
--

DROP TABLE IF EXISTS `ressource`;
CREATE TABLE IF NOT EXISTS `ressource` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `nom` varchar(255) COLLATE utf8_bin NOT NULL,
  `createur` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `date_creation` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nom` (`nom`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Structure de la table `type`
--

DROP TABLE IF EXISTS `type`;
CREATE TABLE IF NOT EXISTS `type` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Structure de la table `type_regle`
--

DROP TABLE IF EXISTS `type_regle`;
CREATE TABLE IF NOT EXISTS `type_regle` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_type` int(10) UNSIGNED NOT NULL,
  `id_regle` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_type` (`id_type`),
  KEY `id_regle` (`id_regle`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `champ`
--
ALTER TABLE `champ`
  ADD CONSTRAINT `champ_ibfk_1` FOREIGN KEY (`ressource_id`) REFERENCES `ressource` (`id`),
  ADD CONSTRAINT `champ_ibfk_2` FOREIGN KEY (`type_id`) REFERENCES `type` (`id`);

--
-- Contraintes pour la table `regle_parametre`
--
ALTER TABLE `regle_parametre`
  ADD CONSTRAINT `regle_parametre_ibfk_1` FOREIGN KEY (`id_regle`) REFERENCES `regle` (`id`),
  ADD CONSTRAINT `regle_parametre_ibfk_2` FOREIGN KEY (`id_parametre`) REFERENCES `parametre` (`id`);

--
-- Contraintes pour la table `type_regle`
--
ALTER TABLE `type_regle`
  ADD CONSTRAINT `type_regle_ibfk_1` FOREIGN KEY (`id_type`) REFERENCES `type` (`id`),
  ADD CONSTRAINT `type_regle_ibfk_2` FOREIGN KEY (`id_regle`) REFERENCES `regle` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
