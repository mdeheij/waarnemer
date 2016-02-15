SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;


CREATE TABLE IF NOT EXISTS `server` (
  `hostID` varchar(100) NOT NULL,
  `ownerID` int(11) NOT NULL,
  `visible` tinyint(1) NOT NULL,
  `authtoken` varchar(256) NOT NULL,
  `hostname` text NOT NULL,
  `identifier` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `serverupdate` (
  `id` int(11) NOT NULL,
  `hostID` varchar(100) NOT NULL,
  `date` bigint(20) NOT NULL,
  `frequency` int(11) NOT NULL,
  `connections` varchar(200) NOT NULL,
  `cpucores` varchar(200) NOT NULL,
  `cpufreq` varchar(200) NOT NULL,
  `cpuname` varchar(200) NOT NULL,
  `diskarray` varchar(200) NOT NULL,
  `disktotal` varchar(200) NOT NULL,
  `diskusage` varchar(200) NOT NULL,
  `filehandles` varchar(200) NOT NULL,
  `filehandleslimit` varchar(200) NOT NULL,
  `hostname` varchar(200) NOT NULL,
  `hostnameshort` varchar(200) NOT NULL,
  `ssid` varchar(60) NOT NULL,
  `ipv4` varchar(200) NOT NULL,
  `ipv4public` varchar(15) NOT NULL,
  `ipv6` varchar(200) NOT NULL,
  `load` varchar(200) NOT NULL,
  `loadcpu` varchar(200) NOT NULL,
  `loadio` varchar(200) NOT NULL,
  `nic` varchar(200) NOT NULL,
  `osarch` varchar(200) NOT NULL,
  `oskernel` varchar(200) NOT NULL,
  `osname` varchar(200) NOT NULL,
  `ping` varchar(200) NOT NULL,
  `packages` int(11) NOT NULL,
  `processes` varchar(200) NOT NULL,
  `processesarray` text NOT NULL,
  `ramtotal` varchar(200) NOT NULL,
  `ramusage` varchar(200) NOT NULL,
  `rx` varchar(200) NOT NULL,
  `rxdiff` varchar(200) NOT NULL,
  `sessions` varchar(200) NOT NULL,
  `swaptotal` varchar(200) NOT NULL,
  `swapusage` varchar(200) NOT NULL,
  `tx` varchar(200) NOT NULL,
  `txdiff` varchar(200) NOT NULL,
  `uptime` varchar(200) NOT NULL,
  `dockerinstalled` int(11) NOT NULL,
  `dockerps` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL,
  `username` varchar(40) NOT NULL,
  `password` varchar(512) NOT NULL,
  `fullname` varchar(70) NOT NULL,
  `role` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


ALTER TABLE `server`
  ADD PRIMARY KEY (`hostID`);

ALTER TABLE `serverupdate`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `serverupdate`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

