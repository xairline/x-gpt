import React from 'react';
import DeckGL from '@deck.gl/react/typed';
import 'mapbox-gl/dist/mapbox-gl.css';
import {useObserver} from 'mobx-react-lite';
import Map from 'react-map-gl';
import {useStores} from "../stores";
// import {inFlowColors, outFlowColors} from '../../../store/flight-log';
// import {useStores} from 'apps/xws/src/store';
// Set your mapbox access token here
const MAPBOX_ACCESS_TOKEN =
    'pk.eyJ1IjoieGFpcmxpbmUiLCJhIjoiY2xkOGE0eHY2MDExZzNvbnh6NG0zYjdkeSJ9.DBehpQbCB9Sjws8OH7I69A';

export interface MapArchProps {
}

const TERRAIN_IMAGE = `https://api.mapbox.com/v4/mapbox.terrain-rgb/{z}/{x}/{y}.png?access_token=${MAPBOX_ACCESS_TOKEN}`;
const SURFACE_IMAGE = `https://api.mapbox.com/v4/mapbox.satellite/{z}/{x}/{y}@2x.png?access_token=${MAPBOX_ACCESS_TOKEN}`;

// https://docs.mapbox.com/help/troubleshooting/access-elevation-data/#mapbox-terrain-rgb
// Note - the elevation rendered by this example is greatly exagerated!
const ELEVATION_DECODER = {
    rScaler: 6553.6,
    gScaler: 25.6,
    bScaler: 0.1,
    offset: -10000
};

// DeckGL react component
export function MapArch2(props: MapArchProps) {
    const {FlightLogStore} = useStores();

    return useObserver(() => {
        function getTooltip({object}: any) {
            if (!object || !object.item) {
                return null;
            }

            const info = object.item;
            return `${info.AircraftDisplayName}
    DEP: ${info.DepartureFlightInfo?.AirportId} - ${info.DepartureFlightInfo?.AirportName}
    ARR: ${info.ArrivalFlightInfo?.AirportId} - ${info.ArrivalFlightInfo?.AirportName}`;
        }

        const layers: any[] = [
            // new ArcLayer({
            //     id: 'arc',
            //     data: FlightLogStore.mapDataSet.arch as any,
            //     getSourcePosition: (d) => d.source,
            //     getTargetPosition: (d) => d.target,
            //     getSourceColor: (d) =>
            //         (d.gain > 0 ? inFlowColors : outFlowColors)[d.quantile] as any,
            //     getTargetColor: (d) =>
            //         (d.gain > 0 ? outFlowColors : inFlowColors)[d.quantile] as any,
            //     getWidth: 3,
            // }),
            // new TerrainLayer({
            //     id: 'terrain',
            //     minZoom: 0,
            //     maxZoom: 23,
            //     strategy: 'no-overlap',
            //     elevationDecoder: ELEVATION_DECODER,
            //     elevationData: TERRAIN_IMAGE,
            //     texture: SURFACE_IMAGE,
            //     wireframe: false,
            //     color: [255, 255, 255]
            // })
        ];
        const INITIAL_VIEW_STATE = {
            // longitude:
            //   FlightLogStore.mapDataSet?.arch &&
            //   FlightLogStore.mapDataSet?.arch?.length > 0
            //     ? FlightLogStore.mapDataSet?.arch[0]?.source[0]
            //     : -75.6692,
            // latitude:
            //   FlightLogStore.mapDataSet?.arch &&
            //   FlightLogStore.mapDataSet?.arch?.length > 0
            //     ? FlightLogStore.mapDataSet?.arch[0]?.source[1]
            //     : 45.3192,
            longitude: -75.6692,
            latitude: 45.3192,
            zoom: 9,
            pitch: 53,
            bearing: 0,
        };
        return (
            <DeckGL
                initialViewState={INITIAL_VIEW_STATE}
                controller={true}
                layers={layers}
                height={'100%'}
                getTooltip={getTooltip}
            >
                <Map
                    mapStyle="mapbox://styles/mapbox/satellite-streets-v12"
                    mapboxAccessToken={MAPBOX_ACCESS_TOKEN}
                >
                </Map>
            </DeckGL>
        );
    });
}

export default MapArch2;
