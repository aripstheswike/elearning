-- phpMyAdmin SQL Dump
-- version 4.8.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 10, 2018 at 06:22 AM
-- Server version: 10.1.32-MariaDB
-- PHP Version: 7.2.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `be_vendor`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `insert_vendor_tran` (IN `nama` VARCHAR(100), IN `no_hp` VARCHAR(50), IN `email` VARCHAR(100), IN `alamat` TEXT, IN `nama_brand` VARCHAR(100), IN `kota` VARCHAR(50), IN `provinsi` VARCHAR(50), IN `kodepos` VARCHAR(50), IN `website` VARCHAR(100), IN `facebook` VARCHAR(100), IN `instagram` VARCHAR(100), IN `marketplace` VARCHAR(100), IN `kategori` VARCHAR(100), IN `pic1` VARCHAR(100), IN `pic2` VARCHAR(100), IN `pic3` VARCHAR(100))  NO SQL
BEGIN

 DECLARE errno INT;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
    GET CURRENT DIAGNOSTICS CONDITION 1 errno = MYSQL_ERRNO;
    SELECT errno AS MYSQL_ERROR;
    ROLLBACK;
    END;

SELECT SUBSTRING(id,8) sufix FROM `vendor`
WHERE DATE_FORMAT(tanggal,'%Y-%m-%d')=CURDATE()
ORDER BY SUFIX DESC LIMIT 1 INTO @LAST_SUFIX;

SELECT CONCAT("3",DATE_FORMAT(NOW(),'%y%m%d'),IF(@LAST_SUFIX IS NULL,'01',LPAD(@LAST_SUFIX+1, 2, '0'))) INTO @NEW_ID;

START TRANSACTION;

INSERT INTO `vendor`(id, nama, no_hp, email, alamat, nama_brand, kota, provinsi, kodepos, website, facebook, instagram, marketplace, kategori) VALUE(@NEW_ID, nama, no_hp, email, alamat, nama_brand, kota, provinsi, kodepos, website, facebook, instagram, marketplace, kategori );

INSERT INTO `picture`(owner, tipe, path_url) VALUE(@NEW_ID, "vendorpic", pic1);

IF (pic2 IS NOT NULL AND pic2 !='') THEN

INSERT INTO `picture`(owner, tipe, path_url) VALUE(@NEW_ID, "vendorpic", pic2);
END IF;

IF (pic3 IS NOT NULL AND pic3 !='') THEN

INSERT INTO `picture`(owner, tipe, path_url) VALUE(@NEW_ID, "vendorpic", pic3);
END IF;

COMMIT WORK;

SELECT @NEW_ID new_id;

END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `picture`
--

CREATE TABLE `picture` (
  `id` int(11) NOT NULL,
  `owner` varchar(30) NOT NULL,
  `tipe` varchar(30) NOT NULL,
  `path_url` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `vendor`
--

CREATE TABLE `vendor` (
  `id` varchar(30) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `no_hp` varchar(30) NOT NULL,
  `email` varchar(100) NOT NULL,
  `alamat` text NOT NULL,
  `nama_brand` varchar(100) NOT NULL,
  `kota` varchar(100) NOT NULL,
  `provinsi` varchar(30) NOT NULL,
  `kodepos` varchar(20) NOT NULL,
  `website` varchar(100) DEFAULT NULL,
  `facebook` varchar(100) DEFAULT NULL,
  `instagram` varchar(100) DEFAULT NULL,
  `marketplace` varchar(100) DEFAULT NULL,
  `kategori` varchar(100) NOT NULL,
  `tanggal` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` varchar(40) DEFAULT 'waiting',
  `password` varchar(300) DEFAULT NULL,
  `role` tinyint(4) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `vendor`
--

INSERT INTO `vendor` (`id`, `nama`, `no_hp`, `email`, `alamat`, `nama_brand`, `kota`, `provinsi`, `kodepos`, `website`, `facebook`, `instagram`, `marketplace`, `kategori`, `tanggal`, `status`, `password`, `role`) VALUES
('123456', 'super admin', '', 'admin', '', '', '', '', '', '', '', '', '', '', '2018-09-04 03:48:00', 'waiting', '4376ef317c9b749e56d776029eac96d41db2aa105539595e80aa45146c47a35651f8619aa304520c13ea81190257e17c374748308c5de7d8f13daee14f34983b', 7);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `picture`
--
ALTER TABLE `picture`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `vendor`
--
ALTER TABLE `vendor`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `picture`
--
ALTER TABLE `picture`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=120;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
