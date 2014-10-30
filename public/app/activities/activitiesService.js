angular.module('tima').factory('activitiesService', ['$http', '$q', '$filter', function($http, $q, $filter) {

    function removeDeletedActivities(source, dest) {
        dest.forEach(function(activity) {
            var actFound = $filter('filter')(source, {id: activity.id}, true);
            if (!actFound.length) {
                var index = dest.indexOf(activity);
                dest.splice(index, 1);
            }
        });
    }

    function mergeActivities(source, dest) {
        source.forEach(function(activity) {
            var actFound = $filter('filter')(dest, {id: activity.id}, true);
            if (actFound.length) {
                actFound[0].duration = activity.duration;
            } else {
                dest.push(activity);
            }
        });
    }

    function getTotalDuration(activities) {
        var totalDuration = 0;
        activities.forEach(function(activity) {
            totalDuration += activity.duration;
        });
        return totalDuration;
    }

    var service = {
        refreshActivities: function(day, activities) {
            var deferred = $q.defer();

            $http.get('/activities/' + day)
            .success(function(data, status, headers, config) {
                removeDeletedActivities(data, activities);
                mergeActivities(data, activities);
                var totalDuration = getTotalDuration(activities);

                deferred.resolve({
                    totalDuration: totalDuration
                });
            })
            .error(function(data, status, header, config){
                // todo: error handling
                deferred.reject(data, status);
            });

            return deferred.promise;
        },

        saveActivity: function(day, text, hours, minutes) {
            var deferred = $q.defer();
            var activity = {
                day: moment(day, 'YYYY-MM-DD').format('YYYY-MM-DD[T]00:00:00.000[Z]'),
                text: text,
                duration: hours * 60 + minutes
            };

            $http.post('/activities', activity)
            .success(function(data, status, headers, config) {
                deferred.resolve();
            })
            .error(function(data, status, header, config) {
                // todo: error handling
                deferred.reject(data, status);
            });

            return deferred.promise;
        },

        deleteActivity: function(id) {
            var deferred = $q.defer();

            $http.delete('/activities/' + id)
            .success(function() {
                deferred.resolve();
            })
            .error(function(data, status) {
                // todo: error handling
                deferred.reject(data, status);
            });

            return deferred.promise;
        }
    };

    return service;
}]);
