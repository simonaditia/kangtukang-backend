-- phpMyAdmin SQL Dump
-- version 5.1.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 28, 2023 at 03:55 AM
-- Server version: 8.0.28
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kangtukang_db_v3`
--

-- --------------------------------------------------------

--
-- Table structure for table `entries`
--

CREATE TABLE `entries` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `content` text,
  `user_id` bigint UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_customer` longtext,
  `id_tukang` longtext,
  `detail_perbaikan` text,
  `waktu_perbaikan` longtext,
  `status` varchar(191) DEFAULT 'Menunggu Konfirmasi',
  `alamat` longtext,
  `customer_name` longtext,
  `tukang_name` longtext,
  `kategori_tukang` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `created_at`, `updated_at`, `deleted_at`, `id_customer`, `id_tukang`, `detail_perbaikan`, `waktu_perbaikan`, `status`, `alamat`, `customer_name`, `tukang_name`, `kategori_tukang`) VALUES
(1, '2023-05-03 13:58:33.235', '2023-06-28 09:28:51.522', NULL, '1', '4', 'Dapurnya perlu di renovasi', '2023-15-05 08:30:00', 'Sedang dikerjakan 11111', 'ini alamat', '', '', ''),
(2, '2023-05-03 15:35:03.013', '2023-05-03 15:35:03.013', NULL, '1', '5', 'Garasinya harus dicat', '2023-04-05 13:44:00', 'Menunggu Konfirmasi', 'ini alamatnya', NULL, NULL, NULL),
(3, '2023-05-03 15:37:13.180', '2023-05-03 15:37:13.180', NULL, '1', '6', 'Plafonnya harus dibetulkan karna sudah kropos', '2023-04-05 13:44:00', 'Sedang Berlangsung', 'ini alamatnya ya', NULL, NULL, NULL),
(4, '2023-05-03 15:37:30.365', '2023-05-03 15:37:30.365', NULL, '1', '6', 'Plafonnya harus dibetulkan karna sudah kropos2', '2023-04-05 13:44:00', 'Selesai', 'ini alamatnya ya2', NULL, NULL, NULL),
(5, '2023-05-03 15:55:39.315', '2023-05-03 15:55:39.315', NULL, '1', '6', 'Plafonnya harus dibetulkan karna sudah kropos3', '2023-04-05 13:44:00', 'Dibatalkan', 'ini alamatnya ya3', NULL, NULL, NULL),
(6, '2023-05-03 16:07:56.510', '2023-05-03 16:07:56.510', NULL, '2', '4', 'Renovasi Kamar mandi saya yang bagus dengan cat warna kuning', '2023-10-05 08:00:00', 'Menunggu Konfirmasi', 'ini alamatnya', NULL, NULL, NULL),
(7, '2023-05-04 12:32:10.186', '2023-05-04 12:32:10.186', NULL, '2', '3', 'Tolong cat dengan warna emas', '2023-05-05 13:44:00', 'Ditolak', 'ini alamatnya', '', '', NULL),
(8, '2023-06-26 01:57:20.290', '2023-06-26 01:57:20.290', NULL, '2', '3', 'Tolong cat dengan warna emas', '2023-05-05 13:44:00', 'Ditolak', 'ini alamatnya', '', '', NULL),
(9, '2023-06-27 21:15:14.148', '2023-06-27 21:15:14.148', NULL, '53', '54', 'Tolong renovasi rumah saya yang bagus', '2023-06-27 13:44:00', 'Menunggu Konfirmasi', 'ini alamatnya', '', '', NULL),
(10, '2023-06-27 21:50:15.851', '2023-06-27 21:50:15.851', NULL, '48', '62', 'too fns. sjs', '2023-06-27 13:44:00', 'Menunggu Konfirmasi', 'Jl. Kp. Poncol , Letnan Arsad No.103, RT.006/RW.015, Kayuringin Jaya, Kec. Bekasi Sel., Kota Bks, Jawa Barat 17144, Indonesia', '', '', NULL),
(11, '2023-06-27 21:50:46.068', '2023-06-27 21:50:46.068', NULL, '48', '62', 'too fns. sjsafjbc xvxvb', '2023-06-27 13:44:00', 'Sedang Berlangsung', 'Jl. Kp. Poncol , Letnan Arsad No.103, RT.006/RW.015, Kayuringin Jaya, Kec. Bekasi Sel., Kota Bks, Jawa Barat 17144, Indonesia', '', '', NULL),
(12, '2023-06-27 22:01:12.095', '2023-06-27 22:01:12.095', NULL, '48', '62', 'g kgkhmululul', '2023-06-27 13:44:00', 'Selesai', 'Jl. Kp. Poncol , Letnan Arsad No.103, RT.006/RW.015, Kayuringin Jaya, Kec. Bekasi Sel., Kota Bks, Jawa Barat 17144, Indonesia', '', '', NULL),
(13, '2023-06-27 22:03:06.751', '2023-06-27 22:03:06.751', NULL, '48', '51', 'vlujivil ijbj', '2023-06-27 13:44:00', 'Menunggu Konfirmasi', '', '', '', NULL),
(14, '2023-06-28 08:11:18.144', '2023-06-28 08:11:18.144', NULL, '48', '51', 'dodkjs sks', '2023-06-27 13:44:00', 'Menunggu Konfirmasi', '', '', '', ''),
(15, '2023-06-28 08:42:09.739', '2023-06-28 08:42:09.739', NULL, '48', '51', 'Tolong renovasi rumah saya yang bagus', '2023-06-27 13:44:00', 'Menunggu Konfirmasi', 'ini alamatnya', '', '', ''),
(16, '2023-06-28 08:42:11.407', '2023-06-28 08:42:11.407', NULL, '48', '51', 'Tolong renovasi rumah saya yang bagus', '2023-06-27 13:44:00', 'Menunggu Konfirmasi', 'ini alamatnya', '', '', ''),
(17, '2023-06-28 08:49:49.101', '2023-06-28 09:58:55.328', NULL, '61', '62', 'Tolong renovasi rumah saya yang bagus', '2023-06-27 13:44:00', 'Sedang Berlangsung', 'ini alamatnya', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nama` longtext,
  `no_telp` longtext,
  `email` longtext,
  `password` longtext,
  `alamat` text,
  `role` longtext,
  `kategori` longtext,
  `biaya` longtext,
  `username` varchar(255) DEFAULT NULL,
  `latitude` double DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `distance` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama`, `no_telp`, `email`, `password`, `alamat`, `role`, `kategori`, `biaya`, `username`, `latitude`, `longitude`, `distance`) VALUES
(1, '2023-05-03 12:01:47.099', '2023-05-03 12:01:47.099', NULL, 'tono', '', 'tono@gmail.com', '$2a$10$S5ysC.40cINMIrLEU.n50exsqF0ZS8QdtCNTXGzwDdb5tBnnT.Sg.', '', 'customer', '', NULL, NULL, 0, 0, NULL),
(2, '2023-05-03 12:01:53.916', '2023-05-03 12:01:53.916', NULL, 'jono', '', 'jono@gmail.com', '$2a$10$LGzZRjdBSFQyCOsMmQIDD.mD9QBFg5zFaYaGMKeCA9dBPlovA3QnS', '', 'customer', '', NULL, NULL, 0, 0, NULL),
(3, '2023-05-03 12:01:58.749', '2023-05-03 12:01:58.749', NULL, 'kiki3', '0814384343', 'kiki3@gmail.com', '$2a$10$9m6J24JYroY40lLJ8TJMGui7QS/EbfbYtEbnFYd9WD/CKPzwYH.IK', 'Jl.nangka3', 'tukang', 'Renovasi', NULL, NULL, 0, 0, NULL),
(4, '2023-05-03 12:02:04.941', '2023-05-03 12:02:04.941', NULL, 'nani2', '0814384342', 'nani2@gmail.com', '$2a$10$K1qoXStgxmCfkEvK0YRCtedRHujHQpLNpSq8Qo5PueRnR9oB.Px8S', 'Jl.nangka2', 'tukang', 'Renovasi', NULL, NULL, 0, 0, NULL),
(5, '2023-05-03 12:06:24.571', '2023-05-03 12:06:24.571', NULL, 'toni', '08143843423', 'toni@gmail.com', '$2a$10$2ey54PwKdbLJPNi9MtukFunHK18CPkrcP7u47ieVkRxTHeSkEWPxe', 'Jl.durian jaya', 'tukang', 'Renovasi', NULL, NULL, 0, 0, NULL),
(6, '2023-05-03 12:06:42.168', '2023-05-03 12:06:42.168', NULL, 'jeki', '08143843113', 'jeki@gmail.com', '123', 'Jl.apel jaya', 'tukang', 'Plafon', NULL, NULL, 0, 0, NULL),
(7, '2023-05-10 12:18:57.292', '2023-05-10 12:18:57.292', NULL, 'oo', '', 'oo@gmail.com', '$2a$10$rDqj1e/W3wLfR/DMusLYM.P6CQ.ORwHMtx/XEqx5FnW1F95qilV3.', '', 'customer', '', '', NULL, 0, 0, NULL),
(8, '2023-05-10 12:35:38.475', '2023-05-10 12:35:38.475', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$bKbqIUGLKehhNZKS5aD6ieqxy77zC270e/b7VcEs24S9GNIpacMS.', '', 'customer', '', '', NULL, 0, 0, NULL),
(9, '2023-05-10 12:35:47.824', '2023-05-10 12:35:47.824', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$tCvUt2EuE0UXbb/vDan0auandP6/ZJYiolSDXYimgL293ncxq9Qau', '', 'customer', '', '', NULL, 0, 0, NULL),
(10, '2023-05-10 12:40:01.693', '2023-05-10 12:40:01.693', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$9iyae2zJV4e0gt4Mdz1/4e8UREd7LoTMuP9LQfWdczMtdYjZF9g8e', '', 'customer', '', '', NULL, 0, 0, NULL),
(11, '2023-05-10 12:40:56.896', '2023-05-10 12:40:56.896', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$WCoPS7CZ87PqAifeQbkhY.IMKNskrK9qENGwNZRrcP/3z2A7QXp.O', '', 'customer', '', '', NULL, 0, 0, NULL),
(12, '2023-05-10 12:40:59.846', '2023-05-10 12:40:59.846', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$UURXFT5aDzcAF0lW1AFIDOwC2GoyEA/8wA.sSN.519QXNkpSC75GW', '', 'customer', '', '', NULL, 0, 0, NULL),
(13, '2023-05-10 12:41:01.543', '2023-05-10 12:41:01.543', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$ZTKdGUEjyqJd/t/2JAOHTeKcALikxdN1XvF0z0gAdpx9Ik20TIlei', '', 'customer', '', '', NULL, 0, 0, NULL),
(14, '2023-05-10 12:41:03.183', '2023-05-10 12:41:03.183', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$e8AXLJ6L/Ilj6JqknwvnC.IYlbiY4DGliyq0A2.t2txISMdrZ0cXu', '', 'customer', '', '', NULL, 0, 0, NULL),
(15, '2023-05-10 12:44:30.983', '2023-05-10 12:44:30.983', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$OX6FLwx0e9KYSvDQV/eS7Oz3XOeCQgHyf05oKTpbQOnsl/qNATJ1m', '', 'customer', '', '', NULL, 0, 0, NULL),
(16, '2023-05-10 12:46:33.112', '2023-05-10 12:46:33.112', NULL, 'kono', '', 'kono@gmail.com', '$2a$10$EqHRzcFE41k2wYc/4zg8XO7B5y6uPDAUsiTRQe8lb2Fzv.6r6uj3a', '', 'customer', '', '', NULL, 0, 0, NULL),
(17, '2023-05-10 12:47:35.214', '2023-05-10 12:47:35.214', NULL, 'kino', '', 'kino@gmail.com', '$2a$10$wzvcf6DoPfktIeX6X04kD.7kIjFCc8vbC6twbZL91aGMkFn.nc7ga', '', 'customer', '', '', NULL, 0, 0, NULL),
(18, '2023-05-10 12:47:39.101', '2023-05-10 12:47:39.101', NULL, 'kino', '', 'kino@gmail.com', '$2a$10$K5AAMt7oU44K6fRM2k0Fy.Av0nmueVB5v3m6bSNAqbtEMZtp7/qri', '', 'customer', '', '', NULL, 0, 0, NULL),
(19, '2023-05-10 12:51:13.696', '2023-05-10 12:51:13.696', NULL, 'kino', '', 'kino@gmail.com', '$2a$10$ia0gOdrF./9FOq6yFInKMu3pV8SD.yLuo5E4/ZmvP1KBc9FoJpljW', '', 'customer', '', '', NULL, 0, 0, NULL),
(20, '2023-05-10 12:53:42.312', '2023-05-10 12:53:42.312', NULL, 'kino', '', 'kino@gmail.com', '$2a$10$NHtDYP8xjo9SXj0urDoE3uJj9bh7f00NtsmIr9HjA6CSuMsA5QqNa', '', 'customer', '', '', NULL, 0, 0, NULL),
(21, '2023-05-10 13:01:44.822', '2023-05-10 13:01:44.822', NULL, 'lolo', '', 'lolo@gmail.com', '$2a$10$UmsChJ1.SPvT6d0NGVr0TuwjZecM.MFnWlrjezc1SUocP62MNt2lK', '', 'customer', '', '', NULL, 0, 0, NULL),
(22, '2023-05-10 13:07:11.807', '2023-05-10 13:07:11.807', NULL, 'o', '', 'o', '$2a$10$eqCNrYvWWhj/tFOiz8zdcOBifPWWbyMNTIIhI5f/v.2jgfHBx5xLS', '', 'customer', '', '', NULL, 0, 0, NULL),
(23, '2023-05-10 13:12:12.255', '2023-05-10 13:12:12.255', NULL, 'o', '', 'o', '$2a$10$Dv9V4oi13.Kd53dEOubHYOs/2kghfl633cfHoLdjPrpXkIjebyPvq', '', 'customer', '', '', NULL, 0, 0, NULL),
(24, '2023-05-10 13:14:17.104', '2023-05-10 13:14:17.104', NULL, 'ooo', '', 'ooo@gmail.com', '$2a$10$SL0RTf3ATmfcPU2ux0ksiOou0c7FsT4HyFSnyScl1JztwBusMXdQ2', '', 'customer', '', '', NULL, 0, 0, NULL),
(25, '2023-05-10 13:18:16.364', '2023-05-10 13:18:16.364', NULL, 'qo', '', 'qo@gmail.com', '$2a$10$WUGQ.5r28a.eEqLh3qrBFeN0hQJcs9Tgsu/z52sbh9AHGfHN9eMwO', '', 'customer', '', '', NULL, 0, 0, NULL),
(26, '2023-05-10 13:22:35.527', '2023-05-10 13:22:35.527', NULL, 'xo', '', 'xo@gmail.com', '$2a$10$MTA4Qu/Fw4vuNhIPTLXzQOzGOvcKdAO9N30IDuFvRtQlpFLudTEzK', '', 'customer', '', '', NULL, 0, 0, NULL),
(27, '2023-05-10 13:28:32.004', '2023-05-10 13:28:32.004', NULL, 'xoo', '', 'xoo@gmail.com', '$2a$10$HDXM8sI3yzNsAAEE3lFA9uBa9JxiecWM4N.jFCRH2.0cpwDLwT.5.', '', 'customer', '', '', NULL, 0, 0, NULL),
(28, '2023-05-10 13:32:25.677', '2023-05-10 13:32:25.677', NULL, 'io', '', 'io@gmail.com', '$2a$10$bar/TS29w/2yR0/4DJtkHOjryJ92is3vUFFtcwhuQJ70trdvsTOdK', '', 'customer', '', '', NULL, 0, 0, NULL),
(29, '2023-05-10 13:34:26.221', '2023-05-10 13:34:26.221', NULL, 'io2', '', 'io@gmail.com', '$2a$10$pTc5itGv.V9PETtyKBaLnuXdR2xVc.ou891WKcGLAJTT75yCFmYSO', '', 'customer', '', '', NULL, 0, 0, NULL),
(30, '2023-05-10 13:44:02.804', '2023-05-10 13:44:02.804', NULL, 'ii', '', 'ii@gmail.com', '$2a$10$FWg1xksSe55qifPiLMXdSeBLcYeHR71IE9pgKcUgcYSoAM9Gr.G.O', '', 'tukang', '', '', NULL, 0, 0, NULL),
(31, '2023-05-10 14:02:56.158', '2023-05-10 14:02:56.158', NULL, 'ko', '', 'ko@gmail.com', '$2a$10$LQh/RFbrsY3vsQcLjExfm.rZo9M0HnZ2Ac23l10zNrPzZYb/v42.K', '', 'customer', '', '', NULL, 0, 0, NULL),
(32, '2023-05-10 14:03:41.653', '2023-05-10 14:03:41.653', NULL, 'ko', '', 'ko@gmail.com', '$2a$10$SFDbMipUsCPPbB4BkEWxgeTXakDhCjUxO6bjos5eSQQ2IbOr5yTgS', '', 'customer', '', '', NULL, 0, 0, NULL),
(33, '2023-05-10 14:04:19.046', '2023-05-10 14:04:19.046', NULL, ' ', '', '', '$2a$10$KpXXMUD/vT6z/1p9SD8Bx.mYdVwjVtHO107MNDknTpx32wQ2PgF4m', '', 'customer', '', '', NULL, 0, 0, NULL),
(34, '2023-05-10 14:34:05.696', '2023-05-10 14:34:05.696', NULL, 'yo', '', 'yo@gmail.com', '$2a$10$G67ytxMTqw1LBphsgsPX1ODwur4ZHAud0e7nlKnI0w2Ij8OYOvgYW', '', 'customer', '', '', NULL, 0, 0, NULL),
(35, '2023-05-16 11:40:24.307', '2023-05-16 11:40:24.307', NULL, 'ioo', '', 'ioo@gmail.com', '$2a$10$SJ75smCE1.5IbG6ptDS4U.MkueIz718aSZ6dKOAkVwH2RJXI1Xu4i', '', 'tukang', '', '', NULL, 0, 0, NULL),
(36, '2023-05-16 13:59:52.973', '2023-05-16 13:59:52.973', NULL, 'oooo', '', 'oooo@gmail.com', '$2a$10$salKmqKRzVmaOhmtyd12iupw02BqG1ZEXlDwfE5xNMqPqSj669FsK', '', 'customer', '', '', NULL, 0, 0, NULL),
(37, '2023-05-17 10:07:34.884', '2023-05-17 10:07:34.884', NULL, 'xoo', '', 'xoo@gmail.com', '$2a$10$eD3NWB2V4HtiFwGhMMGwWe.vakptRUyi9.6Tctozlc37NFNzjOp2O', '', 'customer', '', '', NULL, 0, 0, NULL),
(38, '2023-06-20 23:26:01.644', '2023-06-20 23:26:01.644', NULL, 'affan panjul', '', 'panjul@gmail.com', '$2a$10$KjftY5UuhHxSCPWciBQ8a./MeHZJA8/Eqrx88v7NZBqTqRkpYdaYi', '', 'tukang', '', '', NULL, 0, 0, NULL),
(39, '2023-06-20 23:32:11.534', '2023-06-20 23:32:11.534', NULL, 'dgs', '', 'sgef', '$2a$10$OKalffzQ0nwozpH99FbJe.IZeAOhK5n.iTWmNYGzLV.ffgDkG4tpC', '', 'tukang', '', '', NULL, 0, 0, NULL),
(40, '2023-06-21 11:30:15.083', '2023-06-21 11:30:15.083', NULL, 'jdmnf', '', 'jdkdn@gmail.com', '$2a$10$rEknipWxAeNl8uBw5a.NkODHst2hVCTImEASV4BjXhnhemoUfJbPC', '', 'customer', '', '', NULL, 0, 0, NULL),
(41, '2023-06-21 11:32:46.009', '2023-06-21 11:32:46.009', NULL, 'ygv', '', 'ggvg', '$2a$10$Y1RINJsIT4h/jfuNGV0xBOSrl.Qfb4Kl/abIWBmlz/qlJ9TZ7LU/.', '', 'customer', '', '', NULL, 0, 0, NULL),
(42, '2023-06-21 11:42:24.009', '2023-06-21 11:42:24.009', NULL, 'aaabfg', '', 'cgrg', '$2a$10$g.s2mqCXdNUoRwCgyXKxoeLBFnEPgMQMhzlb3AHTFCbzMHkxLL4zK', '', 'tukang', '', '', NULL, 0, 0, NULL),
(43, '2023-06-21 11:43:30.210', '2023-06-21 11:43:30.210', NULL, 'gtgt g', '', 'vgvy', '$2a$10$lpQBKenGyAswJR8CYXUs4OGZDRs7WPq/KZgsHDISs2w.obZpb8x0O', '', 'customer', '', '', NULL, 0, 0, NULL),
(44, '2023-06-21 11:43:48.713', '2023-06-21 11:43:48.713', NULL, 'gygyv', '', 'tggt', '$2a$10$f/1k3ZFYo2VLplGovy7ra.DzmHVQ8sqOZ4Z5t1XXtSmxgMiv2EDYy', '', 'tukang', '', '', NULL, 0, 0, NULL),
(45, '2023-06-21 11:50:48.851', '2023-06-21 11:50:48.851', NULL, 'jzzhzisbsj shs su shs', '', 'jzz@gmail.com', '$2a$10$5f55m7DgJ3aoaFyhjBf59Oxv0KkEltmqRw7pCTZvWmUowssc3zyWi', '', 'tukang', '', '', NULL, 0, 0, NULL),
(46, '2023-06-21 11:54:05.174', '2023-06-21 11:54:05.174', NULL, 'bla', '', 'bla', '$2a$10$MteRyufvhpURu..IyWrx8OgcAkphWM.9lQiPmQv1WlBIor6/zNuuO', '', 'customer', '', '', NULL, 0, 0, NULL),
(47, '2023-06-21 11:54:42.564', '2023-06-21 11:54:42.564', NULL, 'blo', '', 'blo', '$2a$10$Fp/XeZFqUA8y.TmR.P532uFIhgbgWJYtj9SASo8GN1L1RHXpJyoGi', '', 'tukang', '', '', NULL, 0, 0, NULL),
(48, '2023-06-23 10:24:21.104', '2023-06-23 10:24:21.104', NULL, 'maslat', '', 'maslat@gmail.com', '$2a$10$iwo2I8aCls8RBGIyq4Ep6eOfb2xC1hHh5LcU20Z35V7cfe5T0TMFm', '', 'customer', '', '', NULL, -6.2345036, 106.9840831, NULL),
(50, '2023-06-23 10:31:46.326', '2023-06-23 10:31:46.326', NULL, 'mislat', '', 'mislat@gmail.com', '$2a$10$yPqiCurdE63q5RpqKXwvsumotdu.KC6iQs3K6Sd0hEkjug8kxenHi', '', 'tukang', 'Cat', '', NULL, -6.2376927, 106.9727969, NULL),
(51, '2023-06-26 01:44:48.716', '2023-06-26 01:44:48.716', NULL, 'mislit', '', 'mislit@gmail.com', '$2a$10$JeyiJHRK7x3Ov.1ADIG5zulI5.LdYOik70yqVYQgKyM2imd/MAwd6', '', 'tukang', 'Cat', '', NULL, -6.239593, 106.984843, NULL),
(52, '2023-06-27 15:32:00.227', '2023-06-27 15:32:00.227', NULL, 'coba lokasi customer', '', 'lok@gmail.com', '$2a$10$I1LIr3VDTq25UOsokmUON.Bki2QL0RtlSycl5vhPjrqOcxkpBhKsS', '', 'customer', '', '', NULL, -6.2344204, 106.9861521, 0),
(53, '2023-06-27 15:41:18.876', '2023-06-27 15:41:18.876', NULL, 'lololo lokasi cus', '', 'lololo@gmail.com', '$2a$10$BMlKV9zJEHtq2mlb69RWY.944n9nK1i6yMU1.fLhnsmGK3sNY80BC', '', 'customer', '', '', NULL, -6.2344189, 106.9861538, 0),
(54, '2023-06-27 15:47:14.262', '2023-06-27 15:47:14.262', NULL, 'lilili', '', 'lili@gmail.com', '$2a$10$ySne.JS8bnF4KB.ff/LuKucZBB9syFXvTxgz5uX/iaGZwpjEKOnzi', '', 'tukang', '', '', NULL, -6.2344199, 106.9861508, 0),
(55, '2023-06-27 17:38:30.022', '2023-06-27 17:38:30.022', NULL, 'soissjJ aisjs', '', 'io@gmail.comp', '$2a$10$HyHczaD9IYExoOh8NfSfe.Ru2v0CYTJt79ruPljNY3OkUMtiHF.Sy', '', 'customer', '', '', NULL, -6.2344165, 106.9861515, 0),
(56, '2023-06-27 17:44:03.477', '2023-06-27 17:44:03.477', NULL, 'ususus she s s', '', 'io@gmail.comj', '$2a$10$iZFth9pe9xqjA/oThV8cDOxO3k7Rv/vJmS3Uj2yUw.H1gXEx3UuO6', '', 'customer', '', '', NULL, -6.2344177, 106.9861511, 0),
(57, '2023-06-27 17:46:14.601', '2023-06-27 17:46:14.601', NULL, 'isjshsb r r rff', '', 'io@gmail.compp', '$2a$10$j4r/bm2u/T/LBD8jMgj/uejMg2jU7YFxHXK3BoN1o04u2LaOvcYzi', '', 'customer', '', '', NULL, -6.2343408, 106.9860481, 0),
(58, '2023-06-27 17:48:49.486', '2023-06-27 17:48:49.486', NULL, 'ksisidndje e. d ', '', 'io@gmail.comppd', '$2a$10$kfwJYyUTVKsEPdGugbt2QediUednIOByvaWV79un5GFR6Z2K4PzVO', '', 'customer', '', '', NULL, -6.2344645, 106.9862649, 0),
(59, '2023-06-27 17:50:44.511', '2023-06-27 17:50:44.511', NULL, 'sldd duhdbdb', '', 'io@gmail.com@', '$2a$10$YqveSdHIGZo/RNTf4iAjXuzsV9LfzKw6vMQZOfldKEG.aqs.kAUp6', '', 'customer', '', '', NULL, -6.2344167, 106.9860344, 0),
(60, '2023-06-27 17:57:17.005', '2023-06-27 17:57:17.005', NULL, 'alla sjendnd ', '', 'tukio@gmail.com', '$2a$10$ztGlueXVb5ADIlFoeJ44rODUDklO4NpLHvg2.gSqsljK3yDDxhOiy', '', 'tukang', 'Renovasi', '', NULL, -6.2344189, 106.9861527, 0),
(61, '2023-06-27 18:52:20.975', '2023-06-27 18:52:20.975', NULL, 'alamat civa', '', 'io@gmail.con', '$2a$10$CGndtmp9QgiCgzTjFPl.oOpsf1ORuZ6F3ah9ndb7N0aYWcfkqlgda', 'Jl. Kp. Poncol , Letnan Arsad No.103, RT.006/RW.015, Kayuringin Jaya, Kec. Bekasi Sel., Kota Bks, Jawa Barat 17144, Indonesia', 'customer', '', '', NULL, -6.2344175, 106.9861481, 0),
(62, '2023-06-27 19:13:16.264', '2023-06-27 19:13:16.264', NULL, 'isjsnnz ss s s', '', 'l', '$2a$10$uo0T5/1T8IuWj4jWnkRlgOtXw1uqC3.1YIVjkpgZDI5rDVWSHDuMu', 'Jl. Kp. Poncol , Letnan Arsad No.103, RT.006/RW.015, Kayuringin Jaya, Kec. Bekasi Sel., Kota Bks, Jawa Barat 17144, Indonesia', 'tukang', 'Renovasi', '', NULL, -6.2344157, 106.9861481, 0),
(63, '2023-06-27 19:14:30.247', '2023-06-27 19:14:30.247', NULL, 'pslskms ejsn sisja', '', 'll', '$2a$10$C4WTl7BqGtp8Ba1Rw.mpoerWqCzyy.XLTY2s2GaMEW2QvFVbSBHaW', 'Jl. Kp. Poncol , Letnan Arsad No.103, RT.006/RW.015, Kayuringin Jaya, Kec. Bekasi Sel., Kota Bks, Jawa Barat 17144, Indonesia', 'tukang', 'Renovasi', '', NULL, -6.234417, 106.9861495, 0),
(64, '2023-06-27 19:13:16.264', '2023-06-27 19:13:16.264', NULL, 'isjsnnz ss s s', '', 'l', '$2a$10$uo0T5/1T8IuWj4jWnkRlgOtXw1uqC3.1YIVjkpgZDI5rDVWSHDuMu', 'Jl. Kp. Poncol , Letnan Arsad No.103, RT.006/RW.015, Kayuringin Jaya, Kec. Bekasi Sel., Kota Bks, Jawa Barat 17144, Indonesia', 'tukang', 'Renovasi', '', NULL, -6.2344157, 106.9861481, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `entries`
--
ALTER TABLE `entries`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_entries_deleted_at` (`deleted_at`),
  ADD KEY `fk_users_entries` (`user_id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_orders_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `entries`
--
ALTER TABLE `entries`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=65;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `entries`
--
ALTER TABLE `entries`
  ADD CONSTRAINT `fk_users_entries` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
