/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

CREATE DATABASE IF NOT EXISTS `elearning` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `elearning`;

CREATE TABLE IF NOT EXISTS `t_kelas` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pengajar` int(11) DEFAULT NULL,
  `namakelas` varchar(32) NOT NULL DEFAULT 'Tidak Bernama',
  `hari` varchar(15) DEFAULT NULL,
  `deskripsi` longtext NOT NULL DEFAULT 'Belum Ada Deskripsi',
  `kuota` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `date_created` datetime DEFAULT current_timestamp(),
  `date_updated` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `pengajar` (`pengajar`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_kelas_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_kelas` int(10) unsigned NOT NULL,
  `id_user` int(10) unsigned NOT NULL,
  `date_created` datetime DEFAULT current_timestamp(),
  `date_updated` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_kelas` (`id_kelas`),
  KEY `id_user` (`id_user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_materi` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_kelas` int(10) NOT NULL,
  `judul_materi` varchar(50) DEFAULT NULL,
  `path` varchar(50) DEFAULT NULL,
  `date_created` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_kelas` (`id_kelas`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_mutasi_absensi` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL DEFAULT 0,
  `id_kelas_user` int(11) NOT NULL DEFAULT 0,
  `tanggal_masuk` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) DEFAULT NULL,
  `email` varchar(32) DEFAULT NULL,
  `tipe` tinyint(1) DEFAULT 0 COMMENT '0. Admin\r\n1. Pengajar\r\n3. Siswa',
  `status` tinyint(1) DEFAULT 1,
  `password` varchar(256) DEFAULT NULL,
  `date_created` datetime DEFAULT current_timestamp(),
  `date_updated` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
