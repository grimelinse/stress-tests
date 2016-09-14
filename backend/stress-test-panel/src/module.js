import {PanelCtrl} from  'app/plugins/sdk';
import angular from 'angular';
import $ from  'jquery';

/**
 * Panel for grafana dashboard controls start, stop and obtain the status of
 * the media servers stress tests.
 * This panel controls  makes calls of remote methods of Redis DB client
 * application API:
 *  - /status  Returns status of tests as JSON object.
 *             Expects object like:
 *               {
 *                 "ID":0,
 *                 "Description":"Test ready. Click start for starts this."
 *               },
 *               where
 *               ID - status id:
 *                 0 - The Stress test application is ready for start.
 *                 1 - The stress test application is running.
 *                 2 - The stress test application on error state.
 *              Description - eny text description of status.
 *
 *  - /start_test  Expects parameters:
 *                   model_count - count of publishers.
 *                   client_count - count of players for every publisher.
 *                 Returns status.
 *  - /stop_test   Returns status.
 *
 */
class StressTestControlPanel extends PanelCtrl {

    /**
     * Constructs new StressTestControlPanel.
     * Just calls constructor of superclass.
     *
     * @param $scope      Components model.
     * @param $injector   Application DI injector.
     */
    constructor($scope, $injector) {
        super($scope, $injector);
        this.$http = $injector.get('$http');
        this.events.on('panel-initialized', this.render.bind(this));
        this.events.on('render', this.onRender.bind(this));
        this.$scope.status = "wait";
        this.$scope.status_id = 2;
        this.$scope.stress_test = {
            server: "rtmp://rtmp_server:1935/live",
            model_count: 1,
            client_count: 1
        };
    }

    /**
     * Panel renders event handler.
     * Load redis client configuration, end then requests status of stress test
     * application.
     */
    onRender() {
        var _that = this;
        this.$http.get("./public/redis/config.txt").then(function(response) {
            console.log("redis url: " + response.data);
            _that.$scope.redis_url = response.data;
            _that.getStatus()
        }, function(response) {
            console.log("ERROR loading redis url!");
        })
    }

    /**
     * Requests status of status of stress test application from redis client.
     */
    getStatus() {
        this.$http.get(this.$scope.redis_url + "/status")
            .then(this.onStatus.bind(this)
            , this.onError.bind(this));
    }

    /**
     * Starts stress test.
     */
    startTest() {
        var xsrf = $.param(this.$scope.stress_test);
        this.$http(
            {
                method: 'POST',
                url: this.$scope.redis_url + '/start_test',
                data: xsrf,
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8'
                }
            }
        ).then(this.onStatus.bind(this), this.onError.bind(this));
    }

    /**
     * Stops stress test.
     */
    stopTest() {
        this.$http.get(this.$scope.redis_url + '/stop_test',
            {},
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                }
            }
        ).then(this.onStatus.bind(this), this.onError.bind(this));
    }

    /**
     * Sets status of panel from redis client response.
     *
     * @param response   JSON response from redis client.
     */
    onStatus(response) {
        this.$scope.status_id = response.data.ID;
        this.$scope.status = response.data.Description;
    }

    /**
     * Process response error.
     * @param response   JSON response from redis client.
     */
    onError(response) {
        this.$scope.status = "Error";
        this.$scope.status_id = 2;
    }
}

/**
 * Sets HTML template.
 *
 * @type {string}
 */
StressTestControlPanel.templateUrl = 'module.html';

export {
    StressTestControlPanel as PanelCtrl
    };

