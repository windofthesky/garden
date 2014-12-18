// Code generated by protoc-gen-gogo.
// source: info.proto
// DO NOT EDIT!

package garden

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type InfoRequest struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}

func (m *InfoRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

type InfoResponse struct {
	State            *string                     `protobuf:"bytes,10,opt,name=state" json:"state,omitempty"`
	Events           []string                    `protobuf:"bytes,20,rep,name=events" json:"events,omitempty"`
	HostIp           *string                     `protobuf:"bytes,30,opt,name=host_ip" json:"host_ip,omitempty"`
	ContainerIp      *string                     `protobuf:"bytes,31,opt,name=container_ip" json:"container_ip,omitempty"`
	ContainerPath    *string                     `protobuf:"bytes,32,opt,name=container_path" json:"container_path,omitempty"`
	ExternalIp       *string                     `protobuf:"bytes,33,opt,name=external_ip" json:"external_ip,omitempty"`
	MemoryStat       *InfoResponse_MemoryStat    `protobuf:"bytes,40,opt,name=memory_stat" json:"memory_stat,omitempty"`
	CpuStat          *InfoResponse_CpuStat       `protobuf:"bytes,41,opt,name=cpu_stat" json:"cpu_stat,omitempty"`
	DiskStat         *InfoResponse_DiskStat      `protobuf:"bytes,42,opt,name=disk_stat" json:"disk_stat,omitempty"`
	BandwidthStat    *InfoResponse_BandwidthStat `protobuf:"bytes,43,opt,name=bandwidth_stat" json:"bandwidth_stat,omitempty"`
	ProcessIds       []uint64                    `protobuf:"varint,44,rep,name=process_ids" json:"process_ids,omitempty"`
	Properties       []*Property                 `protobuf:"bytes,45,rep,name=properties" json:"properties,omitempty"`
	MappedPorts      []*InfoResponse_PortMapping `protobuf:"bytes,46,rep,name=mapped_ports" json:"mapped_ports,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}

func (m *InfoResponse) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *InfoResponse) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *InfoResponse) GetHostIp() string {
	if m != nil && m.HostIp != nil {
		return *m.HostIp
	}
	return ""
}

func (m *InfoResponse) GetContainerIp() string {
	if m != nil && m.ContainerIp != nil {
		return *m.ContainerIp
	}
	return ""
}

func (m *InfoResponse) GetContainerPath() string {
	if m != nil && m.ContainerPath != nil {
		return *m.ContainerPath
	}
	return ""
}

func (m *InfoResponse) GetExternalIp() string {
	if m != nil && m.ExternalIp != nil {
		return *m.ExternalIp
	}
	return ""
}

func (m *InfoResponse) GetMemoryStat() *InfoResponse_MemoryStat {
	if m != nil {
		return m.MemoryStat
	}
	return nil
}

func (m *InfoResponse) GetCpuStat() *InfoResponse_CpuStat {
	if m != nil {
		return m.CpuStat
	}
	return nil
}

func (m *InfoResponse) GetDiskStat() *InfoResponse_DiskStat {
	if m != nil {
		return m.DiskStat
	}
	return nil
}

func (m *InfoResponse) GetBandwidthStat() *InfoResponse_BandwidthStat {
	if m != nil {
		return m.BandwidthStat
	}
	return nil
}

func (m *InfoResponse) GetProcessIds() []uint64 {
	if m != nil {
		return m.ProcessIds
	}
	return nil
}

func (m *InfoResponse) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *InfoResponse) GetMappedPorts() []*InfoResponse_PortMapping {
	if m != nil {
		return m.MappedPorts
	}
	return nil
}

type InfoResponse_MemoryStat struct {
	Cache                   *uint64 `protobuf:"varint,1,opt,name=cache" json:"cache,omitempty"`
	Rss                     *uint64 `protobuf:"varint,2,opt,name=rss" json:"rss,omitempty"`
	MappedFile              *uint64 `protobuf:"varint,3,opt,name=mapped_file" json:"mapped_file,omitempty"`
	Pgpgin                  *uint64 `protobuf:"varint,4,opt,name=pgpgin" json:"pgpgin,omitempty"`
	Pgpgout                 *uint64 `protobuf:"varint,5,opt,name=pgpgout" json:"pgpgout,omitempty"`
	Swap                    *uint64 `protobuf:"varint,6,opt,name=swap" json:"swap,omitempty"`
	Pgfault                 *uint64 `protobuf:"varint,7,opt,name=pgfault" json:"pgfault,omitempty"`
	Pgmajfault              *uint64 `protobuf:"varint,8,opt,name=pgmajfault" json:"pgmajfault,omitempty"`
	InactiveAnon            *uint64 `protobuf:"varint,9,opt,name=inactive_anon" json:"inactive_anon,omitempty"`
	ActiveAnon              *uint64 `protobuf:"varint,10,opt,name=active_anon" json:"active_anon,omitempty"`
	InactiveFile            *uint64 `protobuf:"varint,11,opt,name=inactive_file" json:"inactive_file,omitempty"`
	ActiveFile              *uint64 `protobuf:"varint,12,opt,name=active_file" json:"active_file,omitempty"`
	Unevictable             *uint64 `protobuf:"varint,13,opt,name=unevictable" json:"unevictable,omitempty"`
	HierarchicalMemoryLimit *uint64 `protobuf:"varint,14,opt,name=hierarchical_memory_limit" json:"hierarchical_memory_limit,omitempty"`
	HierarchicalMemswLimit  *uint64 `protobuf:"varint,15,opt,name=hierarchical_memsw_limit" json:"hierarchical_memsw_limit,omitempty"`
	TotalCache              *uint64 `protobuf:"varint,16,opt,name=total_cache" json:"total_cache,omitempty"`
	TotalRss                *uint64 `protobuf:"varint,17,opt,name=total_rss" json:"total_rss,omitempty"`
	TotalMappedFile         *uint64 `protobuf:"varint,18,opt,name=total_mapped_file" json:"total_mapped_file,omitempty"`
	TotalPgpgin             *uint64 `protobuf:"varint,19,opt,name=total_pgpgin" json:"total_pgpgin,omitempty"`
	TotalPgpgout            *uint64 `protobuf:"varint,20,opt,name=total_pgpgout" json:"total_pgpgout,omitempty"`
	TotalSwap               *uint64 `protobuf:"varint,21,opt,name=total_swap" json:"total_swap,omitempty"`
	TotalPgfault            *uint64 `protobuf:"varint,22,opt,name=total_pgfault" json:"total_pgfault,omitempty"`
	TotalPgmajfault         *uint64 `protobuf:"varint,23,opt,name=total_pgmajfault" json:"total_pgmajfault,omitempty"`
	TotalInactiveAnon       *uint64 `protobuf:"varint,24,opt,name=total_inactive_anon" json:"total_inactive_anon,omitempty"`
	TotalActiveAnon         *uint64 `protobuf:"varint,25,opt,name=total_active_anon" json:"total_active_anon,omitempty"`
	TotalInactiveFile       *uint64 `protobuf:"varint,26,opt,name=total_inactive_file" json:"total_inactive_file,omitempty"`
	TotalActiveFile         *uint64 `protobuf:"varint,27,opt,name=total_active_file" json:"total_active_file,omitempty"`
	TotalUnevictable        *uint64 `protobuf:"varint,28,opt,name=total_unevictable" json:"total_unevictable,omitempty"`
	XXX_unrecognized        []byte  `json:"-"`
}

func (m *InfoResponse_MemoryStat) Reset()         { *m = InfoResponse_MemoryStat{} }
func (m *InfoResponse_MemoryStat) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_MemoryStat) ProtoMessage()    {}

func (m *InfoResponse_MemoryStat) GetCache() uint64 {
	if m != nil && m.Cache != nil {
		return *m.Cache
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetRss() uint64 {
	if m != nil && m.Rss != nil {
		return *m.Rss
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetMappedFile() uint64 {
	if m != nil && m.MappedFile != nil {
		return *m.MappedFile
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetPgpgin() uint64 {
	if m != nil && m.Pgpgin != nil {
		return *m.Pgpgin
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetPgpgout() uint64 {
	if m != nil && m.Pgpgout != nil {
		return *m.Pgpgout
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetSwap() uint64 {
	if m != nil && m.Swap != nil {
		return *m.Swap
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetPgfault() uint64 {
	if m != nil && m.Pgfault != nil {
		return *m.Pgfault
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetPgmajfault() uint64 {
	if m != nil && m.Pgmajfault != nil {
		return *m.Pgmajfault
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetInactiveAnon() uint64 {
	if m != nil && m.InactiveAnon != nil {
		return *m.InactiveAnon
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetActiveAnon() uint64 {
	if m != nil && m.ActiveAnon != nil {
		return *m.ActiveAnon
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetInactiveFile() uint64 {
	if m != nil && m.InactiveFile != nil {
		return *m.InactiveFile
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetActiveFile() uint64 {
	if m != nil && m.ActiveFile != nil {
		return *m.ActiveFile
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetUnevictable() uint64 {
	if m != nil && m.Unevictable != nil {
		return *m.Unevictable
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetHierarchicalMemoryLimit() uint64 {
	if m != nil && m.HierarchicalMemoryLimit != nil {
		return *m.HierarchicalMemoryLimit
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetHierarchicalMemswLimit() uint64 {
	if m != nil && m.HierarchicalMemswLimit != nil {
		return *m.HierarchicalMemswLimit
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalCache() uint64 {
	if m != nil && m.TotalCache != nil {
		return *m.TotalCache
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalRss() uint64 {
	if m != nil && m.TotalRss != nil {
		return *m.TotalRss
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalMappedFile() uint64 {
	if m != nil && m.TotalMappedFile != nil {
		return *m.TotalMappedFile
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalPgpgin() uint64 {
	if m != nil && m.TotalPgpgin != nil {
		return *m.TotalPgpgin
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalPgpgout() uint64 {
	if m != nil && m.TotalPgpgout != nil {
		return *m.TotalPgpgout
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalSwap() uint64 {
	if m != nil && m.TotalSwap != nil {
		return *m.TotalSwap
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalPgfault() uint64 {
	if m != nil && m.TotalPgfault != nil {
		return *m.TotalPgfault
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalPgmajfault() uint64 {
	if m != nil && m.TotalPgmajfault != nil {
		return *m.TotalPgmajfault
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalInactiveAnon() uint64 {
	if m != nil && m.TotalInactiveAnon != nil {
		return *m.TotalInactiveAnon
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalActiveAnon() uint64 {
	if m != nil && m.TotalActiveAnon != nil {
		return *m.TotalActiveAnon
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalInactiveFile() uint64 {
	if m != nil && m.TotalInactiveFile != nil {
		return *m.TotalInactiveFile
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalActiveFile() uint64 {
	if m != nil && m.TotalActiveFile != nil {
		return *m.TotalActiveFile
	}
	return 0
}

func (m *InfoResponse_MemoryStat) GetTotalUnevictable() uint64 {
	if m != nil && m.TotalUnevictable != nil {
		return *m.TotalUnevictable
	}
	return 0
}

type InfoResponse_CpuStat struct {
	Usage            *uint64 `protobuf:"varint,1,opt,name=usage" json:"usage,omitempty"`
	User             *uint64 `protobuf:"varint,2,opt,name=user" json:"user,omitempty"`
	System           *uint64 `protobuf:"varint,3,opt,name=system" json:"system,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InfoResponse_CpuStat) Reset()         { *m = InfoResponse_CpuStat{} }
func (m *InfoResponse_CpuStat) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_CpuStat) ProtoMessage()    {}

func (m *InfoResponse_CpuStat) GetUsage() uint64 {
	if m != nil && m.Usage != nil {
		return *m.Usage
	}
	return 0
}

func (m *InfoResponse_CpuStat) GetUser() uint64 {
	if m != nil && m.User != nil {
		return *m.User
	}
	return 0
}

func (m *InfoResponse_CpuStat) GetSystem() uint64 {
	if m != nil && m.System != nil {
		return *m.System
	}
	return 0
}

type InfoResponse_DiskStat struct {
	BytesUsed        *uint64 `protobuf:"varint,1,opt,name=bytes_used" json:"bytes_used,omitempty"`
	InodesUsed       *uint64 `protobuf:"varint,2,opt,name=inodes_used" json:"inodes_used,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InfoResponse_DiskStat) Reset()         { *m = InfoResponse_DiskStat{} }
func (m *InfoResponse_DiskStat) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_DiskStat) ProtoMessage()    {}

func (m *InfoResponse_DiskStat) GetBytesUsed() uint64 {
	if m != nil && m.BytesUsed != nil {
		return *m.BytesUsed
	}
	return 0
}

func (m *InfoResponse_DiskStat) GetInodesUsed() uint64 {
	if m != nil && m.InodesUsed != nil {
		return *m.InodesUsed
	}
	return 0
}

type InfoResponse_BandwidthStat struct {
	InRate           *uint64 `protobuf:"varint,1,opt,name=in_rate" json:"in_rate,omitempty"`
	InBurst          *uint64 `protobuf:"varint,2,opt,name=in_burst" json:"in_burst,omitempty"`
	OutRate          *uint64 `protobuf:"varint,3,opt,name=out_rate" json:"out_rate,omitempty"`
	OutBurst         *uint64 `protobuf:"varint,4,opt,name=out_burst" json:"out_burst,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InfoResponse_BandwidthStat) Reset()         { *m = InfoResponse_BandwidthStat{} }
func (m *InfoResponse_BandwidthStat) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_BandwidthStat) ProtoMessage()    {}

func (m *InfoResponse_BandwidthStat) GetInRate() uint64 {
	if m != nil && m.InRate != nil {
		return *m.InRate
	}
	return 0
}

func (m *InfoResponse_BandwidthStat) GetInBurst() uint64 {
	if m != nil && m.InBurst != nil {
		return *m.InBurst
	}
	return 0
}

func (m *InfoResponse_BandwidthStat) GetOutRate() uint64 {
	if m != nil && m.OutRate != nil {
		return *m.OutRate
	}
	return 0
}

func (m *InfoResponse_BandwidthStat) GetOutBurst() uint64 {
	if m != nil && m.OutBurst != nil {
		return *m.OutBurst
	}
	return 0
}

type InfoResponse_PortMapping struct {
	HostPort         *uint32 `protobuf:"varint,1,req,name=host_port" json:"host_port,omitempty"`
	ContainerPort    *uint32 `protobuf:"varint,2,req,name=container_port" json:"container_port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InfoResponse_PortMapping) Reset()         { *m = InfoResponse_PortMapping{} }
func (m *InfoResponse_PortMapping) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_PortMapping) ProtoMessage()    {}

func (m *InfoResponse_PortMapping) GetHostPort() uint32 {
	if m != nil && m.HostPort != nil {
		return *m.HostPort
	}
	return 0
}

func (m *InfoResponse_PortMapping) GetContainerPort() uint32 {
	if m != nil && m.ContainerPort != nil {
		return *m.ContainerPort
	}
	return 0
}

func init() {
}
