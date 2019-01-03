// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package band

import (
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
)

var us_902_928 Band

// US_902_928 is the ID of the US frequency plan
const US_902_928 = "US_902_928"

func init() {
	uplinkChannels := make([]Channel, 0, 72)
	for i := 0; i < 64; i++ {
		uplinkChannels = append(uplinkChannels, Channel{
			Frequency:   uint64(902300000 + 200000*i),
			MinDataRate: 0,
			MaxDataRate: 3,
		})
	}
	for i := 0; i < 8; i++ {
		uplinkChannels = append(uplinkChannels, Channel{
			Frequency:   uint64(903000000 + 1600000*i),
			MinDataRate: 4,
			MaxDataRate: 4,
		})
	}

	downlinkChannels := make([]Channel, 0, 8)
	for i := 0; i < 8; i++ {
		downlinkChannels = append(downlinkChannels, Channel{
			Frequency:   uint64(923300000 + 600000*i),
			MinDataRate: 8,
			MaxDataRate: 13,
		})
	}

	us_902_928 = Band{
		ID: US_902_928,

		MaxUplinkChannels: 72,
		UplinkChannels:    uplinkChannels,

		MaxDownlinkChannels: 8,
		DownlinkChannels:    downlinkChannels,

		BandDutyCycles: []DutyCycle{
			{
				MinFrequency: 902000000,
				MaxFrequency: 928000000,
				Value:        1,
			},
		},

		DataRates: [16]DataRate{
			{Rate: types.DataRate{LoRa: "SF10BW125"}, DefaultMaxSize: maxPayloadSize{19}},
			{Rate: types.DataRate{LoRa: "SF9BW125"}, DefaultMaxSize: maxPayloadSize{61}},
			{Rate: types.DataRate{LoRa: "SF8BW125"}, DefaultMaxSize: maxPayloadSize{133}},
			{Rate: types.DataRate{LoRa: "SF7BW125"}, DefaultMaxSize: maxPayloadSize{250}},
			{Rate: types.DataRate{LoRa: "SF8BW500"}, DefaultMaxSize: maxPayloadSize{250}},
			{}, {}, {}, // RFU
			{Rate: types.DataRate{LoRa: "SF12BW500"}, DefaultMaxSize: maxPayloadSize{41}},
			{Rate: types.DataRate{LoRa: "SF11BW500"}, DefaultMaxSize: maxPayloadSize{117}},
			{Rate: types.DataRate{LoRa: "SF10BW500"}, DefaultMaxSize: maxPayloadSize{230}},
			{Rate: types.DataRate{LoRa: "SF9BW500"}, DefaultMaxSize: maxPayloadSize{230}},
			{Rate: types.DataRate{LoRa: "SF8BW500"}, DefaultMaxSize: maxPayloadSize{230}},
			{Rate: types.DataRate{LoRa: "SF7BW500"}, DefaultMaxSize: maxPayloadSize{230}},
			{}, // Used by LinkADRReq starting from LoRaWAN Regional Parameters 1.1, RFU before
		},
		MaxADRDataRateIndex: 3,

		ReceiveDelay1:    defaultReceiveDelay1,
		ReceiveDelay2:    defaultReceiveDelay2,
		JoinAcceptDelay1: defaultJoinAcceptDelay2,
		JoinAcceptDelay2: defaultJoinAcceptDelay2,
		MaxFCntGap:       defaultMaxFCntGap,
		ADRAckLimit:      defaultADRAckLimit,
		ADRAckDelay:      defaultADRAckDelay,
		MinAckTimeout:    defaultAckTimeout - defaultAckTimeoutMargin,
		MaxAckTimeout:    defaultAckTimeout + defaultAckTimeoutMargin,

		DefaultMaxEIRP: 30,
		TxOffset: func() [16]float32 {
			offset := [16]float32{}
			for i := 0; i < 15; i++ {
				offset[i] = float32(0 - 2*i)
			}
			return offset
		}(),
		MaxTxPowerIndex: 14,

		Rx1Channel: channelIndexModulo(8),
		Rx1DataRate: func(idx ttnpb.DataRateIndex, offset uint32, _ bool) (ttnpb.DataRateIndex, error) {
			if idx > 4 {
				return 0, errDataRateIndexTooHigh.WithAttributes("max", 4)
			}
			if offset > 3 {
				return 0, errDataRateOffsetTooHigh.WithAttributes("max", 3)
			}

			si := int(uint32(idx) + 10 - offset)
			switch {
			case si <= 8:
				return 8, nil
			case si >= 13:
				return 13, nil
			}
			return ttnpb.DataRateIndex(si), nil
		},

		GenerateChMasks: makeGenerateChMask72(true),
		ParseChMask:     parseChMask72,

		ImplementsCFList: true,
		CFListType:       ttnpb.CFListType_CHANNEL_MASKS,

		DefaultRx2Parameters: Rx2Parameters{8, 923300000},

		Beacon: Beacon{
			DataRateIndex:    8,
			CodingRate:       "4/5",
			BroadcastChannel: beaconChannelFromFrequencies(usAuBeaconFrequencies),
			PingSlotChannels: usAuBeaconFrequencies[:],
		},

		regionalParameters1_0:       bandIdentity,
		regionalParameters1_0_1:     bandIdentity,
		regionalParameters1_0_2RevA: usBeacon1_0_2,
		regionalParameters1_0_2RevB: composeSwaps(disableCFList1_0_2, disableChMaskCntl51_0_2),
		regionalParameters1_1RevA:   bandIdentity,
	}
	All[US_902_928] = us_902_928
}
