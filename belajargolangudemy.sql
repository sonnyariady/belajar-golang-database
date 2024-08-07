-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for belajargolangudemy
CREATE DATABASE IF NOT EXISTS `belajargolangudemy` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `belajargolangudemy`;

-- Dumping structure for table belajargolangudemy.karyawan
CREATE TABLE IF NOT EXISTS `karyawan` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Nama` varchar(50) NOT NULL DEFAULT '',
  `Jabatan` varchar(100) NOT NULL DEFAULT '',
  `TanggalLahir` date DEFAULT NULL,
  `Gaji` int DEFAULT '0',
  `Married` bit(1) NOT NULL DEFAULT b'0',
  `Pasangan` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table belajargolangudemy.karyawan: ~12 rows (approximately)
INSERT INTO `karyawan` (`Id`, `Nama`, `Jabatan`, `TanggalLahir`, `Gaji`, `Married`, `Pasangan`) VALUES
	(1, 'Sonny', 'Backend', '1983-01-21', 27, b'1', 'Dyah'),
	(2, 'Yogi Septian', 'QA', '2000-09-14', 24, b'0', NULL),
	(3, 'Wirya', 'Data Engineer', NULL, 26, b'0', NULL),
	(4, 'Dharma', 'Backend', '1994-10-26', NULL, b'1', 'Chintya'),
	(5, 'Herman', 'Accounting', '1987-07-19', 20, b'1', 'Dewi'),
	(6, 'Herman', 'Accounting', '1987-07-19', 15, b'1', 'Henny'),
	(7, 'Nanda', 'Project Manager', '1987-07-19', 15, b'0', NULL),
	(8, 'Nanda', 'Project Manager', NULL, NULL, b'0', NULL),
	(9, 'Dimas', 'Frontend', '2000-03-13', NULL, b'0', NULL),
	(10, 'Teddy', 'Business Analyst', '1987-07-19', 15, b'1', 'Henny'),
	(11, 'Emde', 'QA', '1987-07-19', 15, b'0', NULL),
	(12, 'Alifdaffa', 'PMO', NULL, 17, b'0', 'Nindya');

-- Dumping structure for table belajargolangudemy.mahasiswa
CREATE TABLE IF NOT EXISTS `mahasiswa` (
  `id` varchar(30) NOT NULL,
  `nama` varchar(30) NOT NULL,
  `jurusan` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table belajargolangudemy.mahasiswa: ~3 rows (approximately)
INSERT INTO `mahasiswa` (`id`, `nama`, `jurusan`) VALUES
	('001', 'Sonny', 'Informatika'),
	('002', 'Andi', 'Ilmu Hukum'),
	('003', 'Herman', 'Farmasi');

-- Dumping structure for table belajargolangudemy.useraccount
CREATE TABLE IF NOT EXISTS `useraccount` (
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table belajargolangudemy.useraccount: ~3 rows (approximately)
INSERT INTO `useraccount` (`username`, `password`) VALUES
	('admin', 'password'),
	('user1', 'rahasia'),
	('user2', 'mautauaja');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
