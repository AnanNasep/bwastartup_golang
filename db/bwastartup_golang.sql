-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 03 Jul 2021 pada 16.08
-- Versi server: 10.1.37-MariaDB
-- Versi PHP: 7.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bwastartup_golang`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) NOT NULL,
  `user_id` int(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `short_description` varchar(100) NOT NULL,
  `description` text NOT NULL,
  `goal_amount` int(11) NOT NULL,
  `current_amount` int(11) NOT NULL,
  `perks` text NOT NULL,
  `backer_count` int(11) NOT NULL,
  `slug` varchar(50) NOT NULL,
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaign_image`
--

CREATE TABLE `campaign_image` (
  `id` int(11) NOT NULL,
  `campaign_ig` int(11) NOT NULL,
  `file_name` varchar(50) NOT NULL,
  `is_primary` tinyint(1) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transctions`
--

CREATE TABLE `transctions` (
  `id` int(11) NOT NULL,
  `campaign_ig` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `status` varchar(50) NOT NULL,
  `code` varchar(15) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `occupation` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password_hash` varchar(50) NOT NULL,
  `avatar_file_name` varchar(100) NOT NULL,
  `role` varchar(50) NOT NULL,
  `token` varchar(50) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(1, 'Anan Nasep', 'Programmer', 'nana.septiana.kps@gmail.com', 'Anannasep', 'avatar.jpg', 'Admin', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 'Nasep', 'Tracking', 'nana.septiana.kps@gmail.com', 'Anannasep25', 'avatar.jpg', 'user', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(4, 'Test Simpan', '', '', '', '', '', '', '2021-06-29 23:00:14', '2021-06-29 23:00:14'),
(5, 'Test Simpan2', '', '', '', '', '', '', '2021-06-29 23:01:25', '2021-06-29 23:01:25'),
(6, 'Tes Simpan dari service', 'tukang gorengan', 'contoh@gmail.com', '$2a$04$KhAGT1qujZbz037PjIFxq.j35mbC15UHNvui4DN.HVK', '', 'user', '', '2021-07-03 18:24:11', '2021-07-03 18:24:11'),
(7, 'Tes Simpan dari service2', 'tukang gorengan', 'contoh@gmail.com', '$2a$04$x22jQmBZjjvUMepMVqYw6.6E8ZogZNpUtt0wrj4y6OT', '', 'user', '', '2021-07-03 18:36:54', '2021-07-03 18:36:54'),
(8, 'Anan NAsep Dari Postman', 'Anan NAsep Dari Postman', 'Anan@gmail.com', '$2a$04$fCSN4b4vlG6EamiebXQZqu3h7V.UiHWhi9I7eTuKVBu', '', 'user', '', '2021-07-03 19:05:49', '2021-07-03 19:05:49'),
(9, 'Anan NAsep Dari Postman 2', 'Anan NAsep Dari Postman', 'Anan@gmail.com', '$2a$04$PHjuZAJrBS6HzctC/x8ld.XgwhVlbUSJ2Yf9d/tWeeY', '', 'user', '', '2021-07-03 19:12:57', '2021-07-03 19:12:57'),
(10, 'Anan NAsep Dari Postman 3', 'Anan NAsep Dari Postman', 'Anan@gmail.com', '$2a$04$V35PNztF0pDbhMzjIY21h.NLk0KTNVLi0mGjt5vnwRA', '', 'user', '', '2021-07-03 19:13:30', '2021-07-03 19:13:30'),
(11, 'Anan NAsep Dari Postman 3', 'tracking', 'Anan@gmail.com', '$2a$04$C3QMAoXoS6hKynMNmY2n5eB74Zu3e.jdKtF423tjvKQ', '', 'user', '', '2021-07-03 19:14:30', '2021-07-03 19:14:30'),
(12, 'Anan NAsep Dari Postman response', 'tracking', 'Anan@gmail.com', '$2a$04$gFVgM6Zg8C8uBcEBIy8cmueuH.tuA87.koDcYxvuog8', '', 'user', '', '2021-07-03 19:43:30', '2021-07-03 19:43:30'),
(13, 'Anan NAsep Dari Postman response', 'tracking', 'Anan@gmail.com', '$2a$04$hwAXr.IeQGel3f6LwqcfsuxapDe8/a31bEiFQEIaF7q', '', 'user', '', '2021-07-03 19:47:20', '2021-07-03 19:47:20'),
(14, 'Anan NAsep Dari Postman response 2', 'tracking', 'Anan@gmail.com', '$2a$04$K0p8V1TEEFw00Gvg8hoOS.YO4DTFM7rtCa1yyTrqOwU', '', 'user', '', '2021-07-03 19:47:33', '2021-07-03 19:47:33'),
(15, 'Anan NAsep Dari Postman response 3 + formater', 'tracking', 'Anan@gmail.com', '$2a$04$iJyNYAoA6bEjw354H9TgJ.l.qmDLenbGKDAeW.fT6J.', '', 'user', '', '2021-07-03 20:08:05', '2021-07-03 20:08:05');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `campaign_image`
--
ALTER TABLE `campaign_image`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transctions`
--
ALTER TABLE `transctions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `campaign_image`
--
ALTER TABLE `campaign_image`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `transctions`
--
ALTER TABLE `transctions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
