import React from 'react';
import DeckGL from '@deck.gl/react/typed';
import 'mapbox-gl/dist/mapbox-gl.css';
import {useObserver} from 'mobx-react-lite';
import Map from 'react-map-gl';
import {useStores} from "../stores";
import {inFlowColors, outFlowColors} from "../stores/FlightLog";
import {ArcLayer} from '@deck.gl/layers/typed';
// Set your mapbox access token here
const MAPBOX_ACCESS_TOKEN =
    'pk.eyJ1IjoieGFpcmxpbmUiLCJhIjoiY2xkOGE0eHY2MDExZzNvbnh6NG0zYjdkeSJ9.DBehpQbCB9Sjws8OH7I69A';

export interface MapArchProps {
}

// DeckGL react component
export function MapArch(props: MapArchProps) {
    const {FlightLogStore} = useStores();

    return useObserver(() => {
        const layers: any[] = [
            new ArcLayer({
                id: 'arc',
                data: FlightLogStore.mapDataSet.arch as any,
                getSourcePosition: (d) => d.source,
                getTargetPosition: (d) => d.target,
                getSourceColor: (d) =>
                    (d.gain > 0 ? inFlowColors : outFlowColors)[d.quantile] as any,
                getTargetColor: (d) =>
                    (d.gain > 0 ? outFlowColors : inFlowColors)[d.quantile] as any,
                getWidth: 3,
            }),
        ];
        const INITIAL_VIEW_STATE = {
            longitude:
                FlightLogStore.mapDataSet?.arch &&
                FlightLogStore.mapDataSet?.arch?.length > 0
                    ? FlightLogStore.mapDataSet?.arch[0]?.source[0]
                    : -75.6692,
            latitude:
                FlightLogStore.mapDataSet?.arch &&
                FlightLogStore.mapDataSet?.arch?.length > 0
                    ? FlightLogStore.mapDataSet?.arch[0]?.source[1]
                    : 45.3192,
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

export default MapArch;
