// (C) Copyright 2021 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceServerProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerProfileRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"boot_order": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"boot_mode": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_mode": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"pxe_boot_policy": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"bios_option": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"consistency_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manage_bios": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"overridden_settings": {
							Optional: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"network": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"function_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"network_uri": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port_id": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Lom 1:1-a",
						},
						"requested_mbps": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "2500",
						},
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"boot": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ethernet_boot_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"boot_volume_source": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"iscsi": {
										Type:     schema.TypeSet,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"chap_level": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"first_boot_target_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"first_boot_target_port": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"initiator_name_source": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"second_boot_target_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"second_boot_target_port": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"ipv4": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gateway": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip_address_source": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"server_hardware_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Bay",
			},
			"hide_unused_flex_nics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"manage_connections": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"serial_number_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
			},
			"wwn_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
			},
			"mac_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
			},
			"firmware": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"firmware_baseline_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"consistency_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_activation_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_schedule_date_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manage_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"firmware_install_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"local_storage": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_local_storage": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"initialize": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"controllers": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_slot": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"drive_write_cache": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"import_configuration": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"initialize": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"predictive_spare_rebuild": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"logical_drives": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accelerator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bootable": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"drive_number": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"drive_technology": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"num_physical_drives": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"num_spare_drives": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"raid_level": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sas_logical_jbod_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"logical_jbod": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_slot": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"drive_max_size_gb": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"drive_min_size_gb": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"drive_technology": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"erase_data": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"num_physical_drives": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"persistent": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"sas_logical_jbod_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"san_storage": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_os_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manage_san_storage": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"server_hardware_type_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_hardware_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			// schema for ov.SanStorage.VolumeAttachments
			"volume_attachments": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"lun": {
							Type:     schema.TypeString,
							Required: true,
						},
						"lun_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"boot_volume_priority": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"permanent": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"volume_storage_pool_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_storage_system_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_shareable": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"volume_description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_provision_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_provisioned_capacity_bytes": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"storage_paths": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"storage_target_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"target_selector": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"connection_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"targets": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"tcp_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ServerProfileV9",
			},
			"hardware_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_connection": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ilo_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardware_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_slot_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceServerProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfile, err := config.ovClient.GetProfileByName(d.Get("name").(string))
	if err != nil || serverProfile.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(d.Get("name").(string))

	serverHardware, err := config.ovClient.GetServerHardwareByUri(serverProfile.ServerHardwareURI)
	if err != nil {
		return err
	}

	d.Set("hardware_uri", serverHardware.URI.String())
	d.Set("ilo_ip", serverHardware.GetIloIPAddress())
	d.Set("serial_number", serverProfile.SerialNumber.String())

	if val, ok := d.GetOk("public_connection"); ok {
		publicConnection, err := serverProfile.GetConnectionByName(val.(string))
		if err != nil {
			return err
		}
		d.Set("public_mac", publicConnection.MAC)
		d.Set("public_slot_id", publicConnection.ID)
	}

	enclosureGroup, err := config.ovClient.GetEnclosureGroupByUri(serverProfile.EnclosureGroupURI)
	if err != nil {
		return err
	}
	d.Set("enclosure_group", enclosureGroup.Name)

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByUri(serverProfile.ServerHardwareTypeURI)
	if err != nil {
		return err
	}
	d.Set("server_hardware_type", serverHardwareType.Name)
	d.Set("affinity", serverProfile.Affinity)
	d.Set("serial_number_type", serverProfile.SerialNumberType)
	d.Set("wwn_type", serverProfile.WWNType)
	d.Set("mac_type", serverProfile.MACType)
	d.Set("hide_unused_flex_nics", serverProfile.HideUnusedFlexNics)

	var connections []ov.Connection
	if len(serverProfile.ConnectionSettings.Connections) != 0 {
		connections = serverProfile.ConnectionSettings.Connections
	}
	if len(connections) != 0 {
		networks := make([]map[string]interface{}, 0, len(connections))
		for _, rawNet := range connections {
			networks = append(networks, map[string]interface{}{
				"name":           rawNet.Name,
				"function_type":  rawNet.FunctionType,
				"network_uri":    rawNet.NetworkURI.String(),
				"port_id":        rawNet.PortID,
				"requested_mbps": rawNet.RequestedMbps,
			})
		}
		d.Set("network", networks)
	}
	overriddenSettings := make([]interface{}, 0, len(serverProfile.Bios.OverriddenSettings))
	for _, overriddenSetting := range serverProfile.Bios.OverriddenSettings {
		overriddenSettings = append(overriddenSettings, map[string]interface{}{
			"id":    overriddenSetting.ID,
			"value": overriddenSetting.Value,
		})
	}
	if serverProfile.Bios != nil {
		biosOptions := make([]map[string]interface{}, 0, 1)
		biosOptions = append(biosOptions, map[string]interface{}{
			"manage_bios":         serverProfile.Bios.ManageBios,
			"reapply_state":       serverProfile.Bios.ReapplyState,
			"consistency_state":   serverProfile.Bios.ConsistencyState,
			"overridden_settings": overriddenSettings,
		})

		d.Set("bios_option", biosOptions)
	}

	if serverProfile.Boot.ManageBoot {
		bootOrder := make([]interface{}, 0)
		for _, currBoot := range serverProfile.Boot.Order {
			rawBootOrder := d.Get("boot_order").(*schema.Set).List()
			for _, raw := range rawBootOrder {
				if raw == currBoot {
					bootOrder = append(bootOrder, currBoot)
				}
			}
		}
		d.Set("boot_order", bootOrder)
	}
	d.Set("name", serverProfile.Name)
	d.Set("type", serverProfile.Type)
	d.Set("uri", serverProfile.URI.String())
	d.Set("template", serverProfile.ServerProfileTemplateURI.String())
	localStorages := make([]map[string]interface{}, 0, 1)
	localStorages = append(localStorages, map[string]interface{}{
		"reapply_state": serverProfile.LocalStorage.ReapplyState,
	})

	d.Set("local_storage", localStorages)
	controllers := make([]map[string]interface{}, 0, len(serverProfile.LocalStorage.Controllers))
	for _, controller := range serverProfile.LocalStorage.Controllers {

		controllers = append(controllers, map[string]interface{}{
			"device_slot":          controller.DeviceSlot,
			"drive_write_cache":    controller.DriveWriteCache,
			"import_configuration": controller.ImportConfiguration,
			"initialize":           controller.Initialize,
			"mode":                 controller.Mode,
			"predictive_spare_rebuild": controller.PredictiveSpareRebuild,
		})
		logicaldrives := make([]map[string]interface{}, 0, len(controller.LogicalDrives))
		for _, logicaldrive := range controller.LogicalDrives {

			logicaldrives = append(logicaldrives, map[string]interface{}{
				"accelerator":         logicaldrive.Accelerator,
				"bootable":            logicaldrive.Bootable,
				"drive_number":        logicaldrive.DriveNumber,
				"drive_technology":    logicaldrive.DriveTechnology,
				"name":                logicaldrive.Name,
				"num_physical_drives": logicaldrive.NumPhysicalDrives,
				"num_spare_drives":    logicaldrive.NumSpareDrives,
				"raid_level":          logicaldrive.RaidLevel,
				"sas_logical_jbod_id": logicaldrive.SasLogicalJBODId,
			})
		}

		d.Set("logical_drives", logicaldrives)

	}

	d.Set("controllers", controllers)

	saslogicaljbods := make([]map[string]interface{}, 0, len(serverProfile.LocalStorage.SasLogicalJBODs))
	for _, saslogicaljbod := range serverProfile.LocalStorage.SasLogicalJBODs {

		saslogicaljbods = append(saslogicaljbods, map[string]interface{}{
			"description":          saslogicaljbod.Description,
			"device_slot":          saslogicaljbod.DeviceSlot,
			"drive_max_size_gb":    saslogicaljbod.DriveMaxSizeGB,
			"drive_min_size_gb":    saslogicaljbod.DriveMinSizeGB,
			"drive_technology":     saslogicaljbod.DriveTechnology,
			"erase_data":           saslogicaljbod.EraseData,
			"id":                   saslogicaljbod.ID,
			"name":                 saslogicaljbod.Name,
			"num_physical_drives":  saslogicaljbod.NumPhysicalDrives,
			"persistent":           saslogicaljbod.Persistent,
			"sas_logical_jbod_uri": saslogicaljbod.SasLogicalJBODUri,
			"status":               saslogicaljbod.Status,
		})
	}
	d.Set("logical_jbod", saslogicaljbods)

	return nil
}
