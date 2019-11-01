/*
** Zabbix
** Copyright (C) 2001-2019 Zabbix SIA
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License as published by
** the Free Software Foundation; either version 2 of the License, or
** (at your option) any later version.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
**/

package zbxlib

/* cspell:disable */

/*
#cgo CFLAGS: -I${SRCDIR}/../../../../../include

#include "common.h"
#include "sysinfo.h"
#include "module.h"
typedef int (*zbx_agent_check_t)(AGENT_REQUEST *request, AGENT_RESULT *result);

int	SYSTEM_LOCALTIME(AGENT_REQUEST *request, AGENT_RESULT *result);
int	NET_DNS(AGENT_REQUEST *request, AGENT_RESULT *result);
int	NET_DNS_RECORD(AGENT_REQUEST *request, AGENT_RESULT *result);
int	PROC_MEM(AGENT_REQUEST *request, AGENT_RESULT *result);
int	PROC_NUM(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_BOOTTIME(AGENT_REQUEST *request, AGENT_RESULT *result);
int	WEB_PAGE_GET(AGENT_REQUEST *request, AGENT_RESULT *result);
int	WEB_PAGE_PERF(AGENT_REQUEST *request, AGENT_RESULT *result);
int	WEB_PAGE_REGEXP(AGENT_REQUEST *request, AGENT_RESULT *result);
int	NET_TCP_LISTEN(AGENT_REQUEST *request, AGENT_RESULT *result);
int	NET_TCP_PORT(AGENT_REQUEST *request, AGENT_RESULT *result);
int	CHECK_SERVICE(AGENT_REQUEST *request, AGENT_RESULT *result);
int	CHECK_SERVICE_PERF(AGENT_REQUEST *request, AGENT_RESULT *result);
int	NET_UDP_LISTEN(AGENT_REQUEST *request, AGENT_RESULT *result);
int	GET_SENSOR(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_CPU_LOAD(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_CPU_SWITCHES(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_CPU_INTR(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_HW_CHASSIS(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_HW_CPU(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_HW_DEVICES(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_HW_MACADDR(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_SW_OS(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_SW_PACKAGES(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_SWAP_IN(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_SWAP_OUT(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_SWAP_SIZE(AGENT_REQUEST *request, AGENT_RESULT *result);
int	SYSTEM_USERS_NUM(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VFS_DIR_COUNT(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VFS_DIR_SIZE(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VFS_FILE_MD5SUM(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VFS_FILE_REGMATCH(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VFS_FS_DISCOVERY(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VFS_FS_INODE(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VFS_FS_SIZE(AGENT_REQUEST *request, AGENT_RESULT *result);
int	VM_MEMORY_SIZE(AGENT_REQUEST *request, AGENT_RESULT *result);

*/
import "C"

import (
	"unsafe"
)

func resolveMetric(key string) (cfunc unsafe.Pointer) {
	switch key {
	case "system.localtime":
		return unsafe.Pointer(C.SYSTEM_LOCALTIME)
	case "net.dns":
		return unsafe.Pointer(C.NET_DNS)
	case "net.dns.record":
		return unsafe.Pointer(C.NET_DNS_RECORD)
	case "proc.mem":
		return unsafe.Pointer(C.PROC_MEM)
	case "proc.num":
		return unsafe.Pointer(C.PROC_NUM)
	case "system.boottime":
		return unsafe.Pointer(C.SYSTEM_BOOTTIME)
	case "web.page.get":
		return unsafe.Pointer(C.WEB_PAGE_GET)
	case "web.page.perf":
		return unsafe.Pointer(C.WEB_PAGE_PERF)
	case "web.page.regexp":
		return unsafe.Pointer(C.WEB_PAGE_REGEXP)
	case "net.tcp.listen":
		return unsafe.Pointer(C.NET_TCP_LISTEN)
	case "net.tcp.port":
		return unsafe.Pointer(C.NET_TCP_PORT)
	case "net.tcp.service", "net.udp.service":
		return unsafe.Pointer(C.CHECK_SERVICE)
	case "net.tcp.service.perf", "net.udp.service.perf":
		return unsafe.Pointer(C.CHECK_SERVICE_PERF)
	case "net.udp.listen":
		return unsafe.Pointer(C.NET_UDP_LISTEN)
	case "sensor":
		return unsafe.Pointer(C.GET_SENSOR)
	case "system.cpu.load":
		return unsafe.Pointer(C.SYSTEM_CPU_LOAD)
	case "system.cpu.switches":
		return unsafe.Pointer(C.SYSTEM_CPU_SWITCHES)
	case "system.cpu.intr":
		return unsafe.Pointer(C.SYSTEM_CPU_INTR)
	case "system.hw.chassis":
		return unsafe.Pointer(C.SYSTEM_HW_CHASSIS)
	case "system.hw.cpu":
		return unsafe.Pointer(C.SYSTEM_HW_CPU)
	case "system.hw.devices":
		return unsafe.Pointer(C.SYSTEM_HW_DEVICES)
	case "system.hw.macaddr":
		return unsafe.Pointer(C.SYSTEM_HW_MACADDR)
	case "system.sw.os":
		return unsafe.Pointer(C.SYSTEM_SW_OS)
	case "system.sw.packages":
		return unsafe.Pointer(C.SYSTEM_SW_PACKAGES)
	case "system.swap.in":
		return unsafe.Pointer(C.SYSTEM_SWAP_IN)
	case "system.swap.out":
		return unsafe.Pointer(C.SYSTEM_SWAP_OUT)
	case "system.swap.size":
		return unsafe.Pointer(C.SYSTEM_SWAP_SIZE)
	case "system.users.num":
		return unsafe.Pointer(C.SYSTEM_USERS_NUM)
	case "vfs.dir.count":
		return unsafe.Pointer(C.VFS_DIR_COUNT)
	case "vfs.dir.size":
		return unsafe.Pointer(C.VFS_DIR_SIZE)
	case "vfs.file.md5sum":
		return unsafe.Pointer(C.VFS_FILE_MD5SUM)
	case "vfs.file.regmatch":
		return unsafe.Pointer(C.VFS_FILE_REGMATCH)
	case "vfs.fs.discovery":
		return unsafe.Pointer(C.VFS_FS_DISCOVERY)
	case "vfs.fs.inode":
		return unsafe.Pointer(C.VFS_FS_INODE)
	case "vfs.fs.size":
		return unsafe.Pointer(C.VFS_FS_SIZE)
	case "vm.memory.size":
		return unsafe.Pointer(C.VM_MEMORY_SIZE)

	default:
		return
	}
}