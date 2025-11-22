package stac

var featureCollectionDatasetFixture = `{
  "type": "FeatureCollection",
  "features": [
    {
      "stac_version": "1.0.0",
      "stac_extensions": [
        "https://stac-extensions.github.io/file/v2.1.0/schema.json"
      ],
      "type": "Feature",
      "id": "klaeranlagen_mit_finanzkennzahlen_v1_2-AR",
      "bbox": [9.1910985561571, 47.2470222875734, 9.63096748256717, 47.4690406876864],
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [9.191098556, 47.247022288],
            [9.191098556, 47.469040688],
            [9.630967483, 47.469040688],
            [9.630967483, 47.247022288],
            [9.191098556, 47.247022288]
          ]
        ]
      },
      "properties": {
        "datetime": "2025-10-23T14:41:17Z",
        "title": "Abwasserreinigungsanlagen (ARA) mit Finanzkennzahlen - Kanton AR"
      },
      "links": [
        {
          "href": "https://geodienste.ch/stac/collections/klaeranlagen_mit_finanzkennzahlen/items/klaeranlagen_mit_finanzkennzahlen_v1_2-AR",
          "rel": "self",
          "type": "application/geo+json",
          "title": "This document"
        },
        {
          "href": "https://geodienste.ch/stac",
          "rel": "root",
          "type": "application/json"
        },
        {
          "href": "https://geodienste.ch/stac/collections/klaeranlagen_mit_finanzkennzahlen",
          "rel": "parent",
          "type": "application/json"
        },
        {
          "href": "https://geodienste.ch/stac/collections/klaeranlagen_mit_finanzkennzahlen",
          "rel": "collection",
          "type": "application/json"
        }
      ],
      "assets": {
        "interlis": {
          "href": "https://geodienste.ch/downloads/interlis/klaeranlagen_mit_finanzkennzahlen/AR/klaeranlagen_mit_finanzkennzahlen_v1_2_AR_lv95.zip",
          "title": "Abwasserreinigungsanlagen (ARA) mit Finanzkennzahlen - Kanton AR - EPSG:2056 (ITF/XTF ZIP)",
          "type": "application/zip",
          "file:size": 90105,
          "file:checksum": "d51031386666393434633938393366393138",
          "md5": "18ff944c9893f9182ae38d5282c63d6e",
          "roles": [
            "data"
          ]
        },
        "AR_ARA_DB_REP_V1_2.xtf": {
          "href": "https://geodienste.ch/downloads/interlis/klaeranlagen_mit_finanzkennzahlen/AR/AR_ARA_DB_REP_V1_2.xtf",
          "title": "Abwasserreinigungsanlagen (ARA) mit Finanzkennzahlen - Kanton AR - EPSG:2056 (AR_ARA_DB_REP_V1_2.xtf)",
          "type": "application/interlis+xml",
          "file:size": 435359,
          "file:checksum": "d51033363566346162663935313239333062",
          "md5": "365f4abf9512930bbfe28c28e27c397b",
          "roles": [
            "data"
          ]
        },
        "geopackage_zip": {
          "href": "https://geodienste.ch/downloads/geopackage/klaeranlagen_mit_finanzkennzahlen/AR/deu/klaeranlagen_mit_finanzkennzahlen_v1_2_AR_gpkg_lv95.zip",
          "title": "Abwasserreinigungsanlagen (ARA) mit Finanzkennzahlen - Kanton AR - EPSG:2056 (GPKG ZIP)",
          "type": "application/zip",
          "file:size": 76042,
          "file:checksum": "d51038303964653434316362613132356265",
          "md5": "809de441cba125bef9eb46643d52bbce",
          "roles": [
            "data"
          ]
        },
        "shapefile_zip": {
          "href": "https://geodienste.ch/downloads/shapefile/klaeranlagen_mit_finanzkennzahlen/AR/deu/klaeranlagen_mit_finanzkennzahlen_v1_2_AR_shp_lv95.zip",
          "title": "Abwasserreinigungsanlagen (ARA) mit Finanzkennzahlen - Kanton AR - EPSG:2056 (SHP ZIP)",
          "type": "application/zip",
          "file:size": 73552,
          "file:checksum": "d51035356564326131353631363262643163",
          "md5": "55ed2a156162bd1cf0cd53d1b1f1ffe8",
          "roles": [
            "data"
          ]
        },
        "csv_zip": {
          "href": "https://geodienste.ch/downloads/csv/klaeranlagen_mit_finanzkennzahlen/AR/deu/klaeranlagen_mit_finanzkennzahlen_v1_2_AR_csv_lv95.zip",
          "title": "Abwasserreinigungsanlagen (ARA) mit Finanzkennzahlen - Kanton AR - EPSG:2056 (CSV ZIP)",
          "type": "application/zip",
          "file:size": 6650,
          "file:checksum": "d51062383465626336343734613337666230",
          "md5": "b84ebc6474a37fb0c681444731cea11d",
          "roles": [
            "data"
          ]
        }
      }
    }
  ]
}`
