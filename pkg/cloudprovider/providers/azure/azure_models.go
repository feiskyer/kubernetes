/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// VirtualMachineScaleSetVM describes a virtual machine scale set virtual machine.
type VirtualMachineScaleSetVM struct {
	autorest.Response `json:"-"`
	// InstanceID - The virtual machine instance ID.
	InstanceID *string `json:"instanceId,omitempty"`
	// Sku - The virtual machine SKU.
	Sku                                 *Sku `json:"sku,omitempty"`
	*VirtualMachineScaleSetVMProperties `json:"properties,omitempty"`
	// Zones - The virtual machine zones.
	Zones *[]string `json:"zones,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Sku describes a virtual machine scale set sku.
type Sku struct {
	// Name - The sku name.
	Name *string `json:"name,omitempty"`
	// Tier - Specifies the tier of virtual machines in a scale set.<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**
	Tier *string `json:"tier,omitempty"`
	// Capacity - Specifies the number of virtual machines in the scale set.
	Capacity *int64 `json:"capacity,omitempty"`
}

// VirtualMachineScaleSetVMProperties describes the properties of a virtual machine scale set virtual
// machine.
type VirtualMachineScaleSetVMProperties struct {
	// LatestModelApplied - Specifies whether the latest model has been applied to the virtual machine.
	LatestModelApplied *bool `json:"latestModelApplied,omitempty"`
	// VMID - Azure VM unique ID.
	VMID *string `json:"vmId,omitempty"`
	// InstanceView - The virtual machine instance view.
	InstanceView *VirtualMachineScaleSetVMInstanceView `json:"instanceView,omitempty"`
	// HardwareProfile - Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfile `json:"hardwareProfile,omitempty"`
	// StorageProfile - Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfile `json:"storageProfile,omitempty"`
	// OsProfile - Specifies the operating system settings for the virtual machine.
	OsProfile *OSProfile `json:"osProfile,omitempty"`
	// NetworkProfile - Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	// NetworkProfileConfiguration - Specifies the network profile configuration of the virtual machine.
	NetworkProfileConfiguration *VirtualMachineScaleSetVMNetworkProfileConfiguration `json:"networkProfileConfiguration,omitempty"`
	// ProvisioningState - The provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// VirtualMachineScaleSetVMInstanceView the instance view of a virtual machine scale set VM.
type VirtualMachineScaleSetVMInstanceView struct {
	autorest.Response `json:"-"`
	// PlatformUpdateDomain - The Update Domain count.
	PlatformUpdateDomain *int32 `json:"platformUpdateDomain,omitempty"`
	// PlatformFaultDomain - The Fault Domain count.
	PlatformFaultDomain *int32 `json:"platformFaultDomain,omitempty"`
	// Disks - The disks information.
	Disks *[]DiskInstanceView `json:"disks,omitempty"`
	// PlacementGroupID - The placement group in which the VM is running. If the VM is deallocated it will not have a placementGroupId.
	PlacementGroupID *string `json:"placementGroupId,omitempty"`
}

// DiskInstanceView the instance view of the disk.
type DiskInstanceView struct {
	// Name - The disk name.
	Name *string `json:"name,omitempty"`
	// Statuses - The resource status information.
	Statuses *[]InstanceViewStatus `json:"statuses,omitempty"`
}

// InstanceViewStatus instance view status.
type InstanceViewStatus struct {
	// Code - The status code.
	Code *string `json:"code,omitempty"`
	// Level - The level code. Possible values include: 'Info', 'Warning', 'Error'
	Level StatusLevelTypes `json:"level,omitempty"`
	// DisplayStatus - The short localizable label for the status.
	DisplayStatus *string `json:"displayStatus,omitempty"`
	// Message - The detailed status message, including for alerts and error messages.
	Message *string `json:"message,omitempty"`
	// Time - The time of the status.
	Time *date.Time `json:"time,omitempty"`
}

// StatusLevelTypes enumerates the values for status level types.
type StatusLevelTypes string

const (
	// Error ...
	Error StatusLevelTypes = "Error"
	// Info ...
	Info StatusLevelTypes = "Info"
	// Warning ...
	Warning StatusLevelTypes = "Warning"
)

// HardwareProfile specifies the hardware settings for the virtual machine.
type HardwareProfile struct {
	// VMSize - Specifies the size of the virtual machine. For more information about virtual machine sizes, see [Sizes for virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-sizes?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> The available VM sizes depend on region and availability set. For a list of available sizes use these APIs:  <br><br> [List all available virtual machine sizes in an availability set](https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes) <br><br> [List all available virtual machine sizes in a region](https://docs.microsoft.com/rest/api/compute/virtualmachinesizes/list) <br><br> [List all available virtual machine sizes for resizing](https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes). Possible values include: 'VirtualMachineSizeTypesBasicA0', 'VirtualMachineSizeTypesBasicA1', 'VirtualMachineSizeTypesBasicA2', 'VirtualMachineSizeTypesBasicA3', 'VirtualMachineSizeTypesBasicA4', 'VirtualMachineSizeTypesStandardA0', 'VirtualMachineSizeTypesStandardA1', 'VirtualMachineSizeTypesStandardA2', 'VirtualMachineSizeTypesStandardA3', 'VirtualMachineSizeTypesStandardA4', 'VirtualMachineSizeTypesStandardA5', 'VirtualMachineSizeTypesStandardA6', 'VirtualMachineSizeTypesStandardA7', 'VirtualMachineSizeTypesStandardA8', 'VirtualMachineSizeTypesStandardA9', 'VirtualMachineSizeTypesStandardA10', 'VirtualMachineSizeTypesStandardA11', 'VirtualMachineSizeTypesStandardA1V2', 'VirtualMachineSizeTypesStandardA2V2', 'VirtualMachineSizeTypesStandardA4V2', 'VirtualMachineSizeTypesStandardA8V2', 'VirtualMachineSizeTypesStandardA2mV2', 'VirtualMachineSizeTypesStandardA4mV2', 'VirtualMachineSizeTypesStandardA8mV2', 'VirtualMachineSizeTypesStandardB1s', 'VirtualMachineSizeTypesStandardB1ms', 'VirtualMachineSizeTypesStandardB2s', 'VirtualMachineSizeTypesStandardB2ms', 'VirtualMachineSizeTypesStandardB4ms', 'VirtualMachineSizeTypesStandardB8ms', 'VirtualMachineSizeTypesStandardD1', 'VirtualMachineSizeTypesStandardD2', 'VirtualMachineSizeTypesStandardD3', 'VirtualMachineSizeTypesStandardD4', 'VirtualMachineSizeTypesStandardD11', 'VirtualMachineSizeTypesStandardD12', 'VirtualMachineSizeTypesStandardD13', 'VirtualMachineSizeTypesStandardD14', 'VirtualMachineSizeTypesStandardD1V2', 'VirtualMachineSizeTypesStandardD2V2', 'VirtualMachineSizeTypesStandardD3V2', 'VirtualMachineSizeTypesStandardD4V2', 'VirtualMachineSizeTypesStandardD5V2', 'VirtualMachineSizeTypesStandardD2V3', 'VirtualMachineSizeTypesStandardD4V3', 'VirtualMachineSizeTypesStandardD8V3', 'VirtualMachineSizeTypesStandardD16V3', 'VirtualMachineSizeTypesStandardD32V3', 'VirtualMachineSizeTypesStandardD64V3', 'VirtualMachineSizeTypesStandardD2sV3', 'VirtualMachineSizeTypesStandardD4sV3', 'VirtualMachineSizeTypesStandardD8sV3', 'VirtualMachineSizeTypesStandardD16sV3', 'VirtualMachineSizeTypesStandardD32sV3', 'VirtualMachineSizeTypesStandardD64sV3', 'VirtualMachineSizeTypesStandardD11V2', 'VirtualMachineSizeTypesStandardD12V2', 'VirtualMachineSizeTypesStandardD13V2', 'VirtualMachineSizeTypesStandardD14V2', 'VirtualMachineSizeTypesStandardD15V2', 'VirtualMachineSizeTypesStandardDS1', 'VirtualMachineSizeTypesStandardDS2', 'VirtualMachineSizeTypesStandardDS3', 'VirtualMachineSizeTypesStandardDS4', 'VirtualMachineSizeTypesStandardDS11', 'VirtualMachineSizeTypesStandardDS12', 'VirtualMachineSizeTypesStandardDS13', 'VirtualMachineSizeTypesStandardDS14', 'VirtualMachineSizeTypesStandardDS1V2', 'VirtualMachineSizeTypesStandardDS2V2', 'VirtualMachineSizeTypesStandardDS3V2', 'VirtualMachineSizeTypesStandardDS4V2', 'VirtualMachineSizeTypesStandardDS5V2', 'VirtualMachineSizeTypesStandardDS11V2', 'VirtualMachineSizeTypesStandardDS12V2', 'VirtualMachineSizeTypesStandardDS13V2', 'VirtualMachineSizeTypesStandardDS14V2', 'VirtualMachineSizeTypesStandardDS15V2', 'VirtualMachineSizeTypesStandardDS134V2', 'VirtualMachineSizeTypesStandardDS132V2', 'VirtualMachineSizeTypesStandardDS148V2', 'VirtualMachineSizeTypesStandardDS144V2', 'VirtualMachineSizeTypesStandardE2V3', 'VirtualMachineSizeTypesStandardE4V3', 'VirtualMachineSizeTypesStandardE8V3', 'VirtualMachineSizeTypesStandardE16V3', 'VirtualMachineSizeTypesStandardE32V3', 'VirtualMachineSizeTypesStandardE64V3', 'VirtualMachineSizeTypesStandardE2sV3', 'VirtualMachineSizeTypesStandardE4sV3', 'VirtualMachineSizeTypesStandardE8sV3', 'VirtualMachineSizeTypesStandardE16sV3', 'VirtualMachineSizeTypesStandardE32sV3', 'VirtualMachineSizeTypesStandardE64sV3', 'VirtualMachineSizeTypesStandardE3216V3', 'VirtualMachineSizeTypesStandardE328sV3', 'VirtualMachineSizeTypesStandardE6432sV3', 'VirtualMachineSizeTypesStandardE6416sV3', 'VirtualMachineSizeTypesStandardF1', 'VirtualMachineSizeTypesStandardF2', 'VirtualMachineSizeTypesStandardF4', 'VirtualMachineSizeTypesStandardF8', 'VirtualMachineSizeTypesStandardF16', 'VirtualMachineSizeTypesStandardF1s', 'VirtualMachineSizeTypesStandardF2s', 'VirtualMachineSizeTypesStandardF4s', 'VirtualMachineSizeTypesStandardF8s', 'VirtualMachineSizeTypesStandardF16s', 'VirtualMachineSizeTypesStandardF2sV2', 'VirtualMachineSizeTypesStandardF4sV2', 'VirtualMachineSizeTypesStandardF8sV2', 'VirtualMachineSizeTypesStandardF16sV2', 'VirtualMachineSizeTypesStandardF32sV2', 'VirtualMachineSizeTypesStandardF64sV2', 'VirtualMachineSizeTypesStandardF72sV2', 'VirtualMachineSizeTypesStandardG1', 'VirtualMachineSizeTypesStandardG2', 'VirtualMachineSizeTypesStandardG3', 'VirtualMachineSizeTypesStandardG4', 'VirtualMachineSizeTypesStandardG5', 'VirtualMachineSizeTypesStandardGS1', 'VirtualMachineSizeTypesStandardGS2', 'VirtualMachineSizeTypesStandardGS3', 'VirtualMachineSizeTypesStandardGS4', 'VirtualMachineSizeTypesStandardGS5', 'VirtualMachineSizeTypesStandardGS48', 'VirtualMachineSizeTypesStandardGS44', 'VirtualMachineSizeTypesStandardGS516', 'VirtualMachineSizeTypesStandardGS58', 'VirtualMachineSizeTypesStandardH8', 'VirtualMachineSizeTypesStandardH16', 'VirtualMachineSizeTypesStandardH8m', 'VirtualMachineSizeTypesStandardH16m', 'VirtualMachineSizeTypesStandardH16r', 'VirtualMachineSizeTypesStandardH16mr', 'VirtualMachineSizeTypesStandardL4s', 'VirtualMachineSizeTypesStandardL8s', 'VirtualMachineSizeTypesStandardL16s', 'VirtualMachineSizeTypesStandardL32s', 'VirtualMachineSizeTypesStandardM64s', 'VirtualMachineSizeTypesStandardM64ms', 'VirtualMachineSizeTypesStandardM128s', 'VirtualMachineSizeTypesStandardM128ms', 'VirtualMachineSizeTypesStandardM6432ms', 'VirtualMachineSizeTypesStandardM6416ms', 'VirtualMachineSizeTypesStandardM12864ms', 'VirtualMachineSizeTypesStandardM12832ms', 'VirtualMachineSizeTypesStandardNC6', 'VirtualMachineSizeTypesStandardNC12', 'VirtualMachineSizeTypesStandardNC24', 'VirtualMachineSizeTypesStandardNC24r', 'VirtualMachineSizeTypesStandardNC6sV2', 'VirtualMachineSizeTypesStandardNC12sV2', 'VirtualMachineSizeTypesStandardNC24sV2', 'VirtualMachineSizeTypesStandardNC24rsV2', 'VirtualMachineSizeTypesStandardNC6sV3', 'VirtualMachineSizeTypesStandardNC12sV3', 'VirtualMachineSizeTypesStandardNC24sV3', 'VirtualMachineSizeTypesStandardNC24rsV3', 'VirtualMachineSizeTypesStandardND6s', 'VirtualMachineSizeTypesStandardND12s', 'VirtualMachineSizeTypesStandardND24s', 'VirtualMachineSizeTypesStandardND24rs', 'VirtualMachineSizeTypesStandardNV6', 'VirtualMachineSizeTypesStandardNV12', 'VirtualMachineSizeTypesStandardNV24'
	VMSize string `json:"vmSize,omitempty"`
}

// StorageProfile specifies the storage settings for the virtual machine disks.
type StorageProfile struct {
	// ImageReference - Specifies information about the image to use. You can specify information about platform images, marketplace images, or virtual machine images. This element is required when you want to use a platform image, marketplace image, or virtual machine image, but is not used in other creation operations.
	ImageReference *ImageReference `json:"imageReference,omitempty"`
	// OsDisk - Specifies information about the operating system disk used by the virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	OsDisk *OSDisk `json:"osDisk,omitempty"`
	// DataDisks - Specifies the parameters that are used to add a data disk to a virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	DataDisks *[]DataDisk `json:"dataDisks,omitempty"`
}

// ImageReference specifies information about the image to use. You can specify information about platform
// images, marketplace images, or virtual machine images. This element is required when you want to use a
// platform image, marketplace image, or virtual machine image, but is not used in other creation
// operations.
type ImageReference struct {
	// Publisher - The image publisher.
	Publisher *string `json:"publisher,omitempty"`
	// Offer - Specifies the offer of the platform image or marketplace image used to create the virtual machine.
	Offer *string `json:"offer,omitempty"`
	// Sku - The image SKU.
	Sku *string `json:"sku,omitempty"`
	// Version - Specifies the version of the platform image or marketplace image used to create the virtual machine. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available.
	Version *string `json:"version,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// OSDisk specifies information about the operating system disk used by the virtual machine. <br><br> For
// more information about disks, see [About disks and VHDs for Azure virtual
// machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
type OSDisk struct {
	// OsType - This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**. Possible values include: 'Windows', 'Linux'
	OsType OperatingSystemTypes `json:"osType,omitempty"`
	// EncryptionSettings - Specifies the encryption settings for the OS Disk. <br><br> Minimum api-version: 2015-06-15
	EncryptionSettings *DiskEncryptionSettings `json:"encryptionSettings,omitempty"`
	// Name - The disk name.
	Name *string `json:"name,omitempty"`
	// Vhd - The virtual hard disk.
	Vhd *VirtualHardDisk `json:"vhd,omitempty"`
	// Image - The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
	Image *VirtualHardDisk `json:"image,omitempty"`
	// Caching - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**. Possible values include: 'CachingTypesNone', 'CachingTypesReadOnly', 'CachingTypesReadWrite'
	Caching CachingTypes `json:"caching,omitempty"`
	// WriteAcceleratorEnabled - Specifies whether writeAccelerator should be enabled or disabled on the disk.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
	// DiffDiskSettings - Specifies the ephemeral Disk Settings for the operating system disk used by the virtual machine.
	DiffDiskSettings *DiffDiskSettings `json:"diffDiskSettings,omitempty"`
	// CreateOption - Specifies how the virtual machine should be created.<br><br> Possible values are:<br><br> **Attach** \u2013 This value is used when you are using a specialized disk to create the virtual machine.<br><br> **FromImage** \u2013 This value is used when you are using an image to create the virtual machine. If you are using a platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also use the plan element previously described. Possible values include: 'DiskCreateOptionTypesFromImage', 'DiskCreateOptionTypesEmpty', 'DiskCreateOptionTypesAttach'
	CreateOption DiskCreateOptionTypes `json:"createOption,omitempty"`
	// DiskSizeGB - Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the disk in a virtual machine image. <br><br> This value cannot be larger than 1023 GB
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
	// ManagedDisk - The managed disk parameters.
	ManagedDisk *ManagedDiskParameters `json:"managedDisk,omitempty"`
}

// OperatingSystemTypes enumerates the values for operating system types.
type OperatingSystemTypes string

const (
	// Linux ...
	Linux OperatingSystemTypes = "Linux"
	// Windows ...
	Windows OperatingSystemTypes = "Windows"
)

// DiskEncryptionSettings describes a Encryption Settings for a Disk
type DiskEncryptionSettings struct {
	// DiskEncryptionKey - Specifies the location of the disk encryption key, which is a Key Vault Secret.
	DiskEncryptionKey *KeyVaultSecretReference `json:"diskEncryptionKey,omitempty"`
	// KeyEncryptionKey - Specifies the location of the key encryption key in Key Vault.
	KeyEncryptionKey *KeyVaultKeyReference `json:"keyEncryptionKey,omitempty"`
	// Enabled - Specifies whether disk encryption should be enabled on the virtual machine.
	Enabled *bool `json:"enabled,omitempty"`
}

// KeyVaultKeyReference describes a reference to Key Vault Key
type KeyVaultKeyReference struct {
	// KeyURL - The URL referencing a key encryption key in Key Vault.
	KeyURL *string `json:"keyUrl,omitempty"`
	// SourceVault - The relative URL of the Key Vault containing the key.
	SourceVault *SubResource `json:"sourceVault,omitempty"`
}

// KeyVaultSecretReference describes a reference to Key Vault Secret
type KeyVaultSecretReference struct {
	// SecretURL - The URL referencing a secret in a Key Vault.
	SecretURL *string `json:"secretUrl,omitempty"`
	// SourceVault - The relative URL of the Key Vault containing the secret.
	SourceVault *SubResource `json:"sourceVault,omitempty"`
}

// SubResource ...
type SubResource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// VirtualHardDisk describes the uri of a disk.
type VirtualHardDisk struct {
	// URI - Specifies the virtual hard disk's uri.
	URI *string `json:"uri,omitempty"`
}

// CachingTypes enumerates the values for caching types.
type CachingTypes string

const (
	// CachingTypesNone ...
	CachingTypesNone CachingTypes = "None"
	// CachingTypesReadOnly ...
	CachingTypesReadOnly CachingTypes = "ReadOnly"
	// CachingTypesReadWrite ...
	CachingTypesReadWrite CachingTypes = "ReadWrite"
)

// DiffDiskSettings describes the parameters of ephemeral disk settings that can be specified for operating
// system disk. <br><br> NOTE: The ephemeral disk settings can only be specified for managed disk.
type DiffDiskSettings struct {
	// Option - Specifies the ephemeral disk settings for operating system disk. Possible values include: 'Local'
	Option DiffDiskOptions `json:"option,omitempty"`
}

// DiffDiskOptions enumerates the values for diff disk options.
type DiffDiskOptions string

const (
	// Local ...
	Local DiffDiskOptions = "Local"
)

// DiskCreateOptionTypes enumerates the values for disk create option types.
type DiskCreateOptionTypes string

const (
	// DiskCreateOptionTypesAttach ...
	DiskCreateOptionTypesAttach DiskCreateOptionTypes = "Attach"
	// DiskCreateOptionTypesEmpty ...
	DiskCreateOptionTypesEmpty DiskCreateOptionTypes = "Empty"
	// DiskCreateOptionTypesFromImage ...
	DiskCreateOptionTypesFromImage DiskCreateOptionTypes = "FromImage"
)

// ManagedDiskParameters the parameters of a managed disk.
type ManagedDiskParameters struct {
	// StorageAccountType - Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk. Possible values include: 'StorageAccountTypesStandardLRS', 'StorageAccountTypesPremiumLRS', 'StorageAccountTypesStandardSSDLRS', 'StorageAccountTypesUltraSSDLRS'
	StorageAccountType StorageAccountTypes `json:"storageAccountType,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// StorageAccountTypes enumerates the values for storage account types.
type StorageAccountTypes string

const (
	// StorageAccountTypesPremiumLRS ...
	StorageAccountTypesPremiumLRS StorageAccountTypes = "Premium_LRS"
	// StorageAccountTypesStandardLRS ...
	StorageAccountTypesStandardLRS StorageAccountTypes = "Standard_LRS"
	// StorageAccountTypesStandardSSDLRS ...
	StorageAccountTypesStandardSSDLRS StorageAccountTypes = "StandardSSD_LRS"
	// StorageAccountTypesUltraSSDLRS ...
	StorageAccountTypesUltraSSDLRS StorageAccountTypes = "UltraSSD_LRS"
)

// DataDisk describes a data disk.
type DataDisk struct {
	// Lun - Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
	Lun *int32 `json:"lun,omitempty"`
	// Name - The disk name.
	Name *string `json:"name,omitempty"`
	// Vhd - The virtual hard disk.
	Vhd *VirtualHardDisk `json:"vhd,omitempty"`
	// Image - The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
	Image *VirtualHardDisk `json:"image,omitempty"`
	// Caching - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**. Possible values include: 'CachingTypesNone', 'CachingTypesReadOnly', 'CachingTypesReadWrite'
	Caching CachingTypes `json:"caching,omitempty"`
	// WriteAcceleratorEnabled - Specifies whether writeAccelerator should be enabled or disabled on the disk.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
	// CreateOption - Specifies how the virtual machine should be created.<br><br> Possible values are:<br><br> **Attach** \u2013 This value is used when you are using a specialized disk to create the virtual machine.<br><br> **FromImage** \u2013 This value is used when you are using an image to create the virtual machine. If you are using a platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also use the plan element previously described. Possible values include: 'DiskCreateOptionTypesFromImage', 'DiskCreateOptionTypesEmpty', 'DiskCreateOptionTypesAttach'
	CreateOption DiskCreateOptionTypes `json:"createOption,omitempty"`
	// DiskSizeGB - Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the disk in a virtual machine image. <br><br> This value cannot be larger than 1023 GB
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
	// ManagedDisk - The managed disk parameters.
	ManagedDisk *ManagedDiskParameters `json:"managedDisk,omitempty"`
}

// OSProfile specifies the operating system settings for the virtual machine.
type OSProfile struct {
	// ComputerName - Specifies the host OS name of the virtual machine. <br><br> This name cannot be updated after the VM is created. <br><br> **Max-length (Windows):** 15 characters <br><br> **Max-length (Linux):** 64 characters. <br><br> For naming conventions and restrictions see [Azure infrastructure services implementation guidelines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions).
	ComputerName *string `json:"computerName,omitempty"`
}

// NetworkInterfaceReferenceProperties describes a network interface reference properties.
type NetworkInterfaceReferenceProperties struct {
	// Primary - Specifies the primary network interface in case the virtual machine has more than 1 network interface.
	Primary *bool `json:"primary,omitempty"`
}

// NetworkInterfaceReference describes a network interface reference.
type NetworkInterfaceReference struct {
	*NetworkInterfaceReferenceProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// NetworkProfile specifies the network interfaces of the virtual machine.
type NetworkProfile struct {
	// NetworkInterfaces - Specifies the list of resource Ids for the network interfaces associated with the virtual machine.
	NetworkInterfaces *[]NetworkInterfaceReference `json:"networkInterfaces,omitempty"`
}

// VirtualMachineScaleSetVMNetworkProfileConfiguration describes a virtual machine scale set VM network
// profile.
type VirtualMachineScaleSetVMNetworkProfileConfiguration struct {
	// NetworkInterfaceConfigurations - The list of network configurations.
	NetworkInterfaceConfigurations *[]VirtualMachineScaleSetNetworkConfiguration `json:"networkInterfaceConfigurations,omitempty"`
}

// VirtualMachineScaleSetNetworkConfiguration describes a virtual machine scale set network profile's
// network configurations.
type VirtualMachineScaleSetNetworkConfiguration struct {
	// Name - The network configuration name.
	Name                                                  *string `json:"name,omitempty"`
	*VirtualMachineScaleSetNetworkConfigurationProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// VirtualMachineScaleSetNetworkConfigurationProperties describes a virtual machine scale set network
// profile's IP configuration.
type VirtualMachineScaleSetNetworkConfigurationProperties struct {
	// Primary - Specifies the primary network interface in case the virtual machine has more than 1 network interface.
	Primary *bool `json:"primary,omitempty"`
	// EnableAcceleratedNetworking - Specifies whether the network interface is accelerated networking-enabled.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`
	// NetworkSecurityGroup - The network security group.
	NetworkSecurityGroup *SubResource `json:"networkSecurityGroup,omitempty"`
	// DNSSettings - The dns settings to be applied on the network interfaces.
	DNSSettings *VirtualMachineScaleSetNetworkConfigurationDNSSettings `json:"dnsSettings,omitempty"`
	// IPConfigurations - Specifies the IP configurations of the network interface.
	IPConfigurations *[]VirtualMachineScaleSetIPConfiguration `json:"ipConfigurations,omitempty"`
	// EnableIPForwarding - Whether IP forwarding enabled on this NIC.
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`
}

// VirtualMachineScaleSetNetworkConfigurationDNSSettings describes a virtual machines scale sets network
// configuration's DNS settings.
type VirtualMachineScaleSetNetworkConfigurationDNSSettings struct {
	// DNSServers - List of DNS servers IP addresses
	DNSServers *[]string `json:"dnsServers,omitempty"`
}

// VirtualMachineScaleSetIPConfiguration describes a virtual machine scale set network profile's IP
// configuration.
type VirtualMachineScaleSetIPConfiguration struct {
	// Name - The IP configuration name.
	Name                                             *string `json:"name,omitempty"`
	*VirtualMachineScaleSetIPConfigurationProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// VirtualMachineScaleSetIPConfigurationProperties describes a virtual machine scale set network profile's
// IP configuration properties.
type VirtualMachineScaleSetIPConfigurationProperties struct {
	// Subnet - Specifies the identifier of the subnet.
	Subnet *APIEntityReference `json:"subnet,omitempty"`
	// Primary - Specifies the primary network interface in case the virtual machine has more than 1 network interface.
	Primary *bool `json:"primary,omitempty"`
	// PrivateIPAddressVersion - Available from Api-Version 2017-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'. Possible values include: 'IPv4', 'IPv6'
	PrivateIPAddressVersion IPVersion `json:"privateIPAddressVersion,omitempty"`
	// LoadBalancerBackendAddressPools - Specifies an array of references to backend address pools of load balancers. A scale set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the same load balancer.
	LoadBalancerBackendAddressPools *[]SubResource `json:"loadBalancerBackendAddressPools,omitempty"`
	// LoadBalancerInboundNatPools - Specifies an array of references to inbound Nat pools of the load balancers. A scale set can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same load balancer
	LoadBalancerInboundNatPools *[]SubResource `json:"loadBalancerInboundNatPools,omitempty"`
}

// IPVersion enumerates the values for ip version.
type IPVersion string

const (
	// IPv4 ...
	IPv4 IPVersion = "IPv4"
	// IPv6 ...
	IPv6 IPVersion = "IPv6"
)

// APIEntityReference the API entity reference.
type APIEntityReference struct {
	// ID - The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
	ID *string `json:"id,omitempty"`
}
